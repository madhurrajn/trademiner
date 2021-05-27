package excel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"miner/ai/stats"
	"miner/auth"
	"miner/gfin"
	"miner/models"
	"miner/utils"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

type CountHist struct {
	apiCounter int
	ts         int64
}

var MaxThrottle = 55

var ReadApiCounter = map[string]CountHist{"ReadData": CountHist{apiCounter: 0, ts: time.Now().Unix()},
	"WriteData": CountHist{apiCounter: 0, ts: time.Now().Unix()}}

func UpdateCountHist(counthist CountHist, counterName string) {
	fmt.Printf("counthist update\n")
	ReadApiCounter[counterName] = counthist
}

func ThrottleApiRequest(counterName string) bool {
	counthist := ReadApiCounter[counterName]
	fmt.Printf("CountHist %v\n", counthist)
	//defer UpdateCountHist(counthist, counterName)
	if counthist.apiCounter == 0 {
		counthist.apiCounter = 1
		counthist.ts = time.Now().Unix()
		UpdateCountHist(counthist, counterName)
		return false
	}
	counthist.apiCounter = counthist.apiCounter + 1
	if counthist.apiCounter >= MaxThrottle {
		//Throttling Required
		numSecondsLeft := time.Now().Unix() - counthist.ts
		log.Printf("Num Seconds Left %v", numSecondsLeft)
		fmt.Printf("Num Seconds Left %v", numSecondsLeft)
		//time.Sleep(time.Duration(numSecondsLeft) * time.Second)
		time.Sleep(10 * time.Second)
		counthist.apiCounter = 0
		counthist.ts = time.Now().Unix()
		UpdateCountHist(counthist, counterName)
		return true
	}
	numSeconds := time.Now().Unix() - counthist.ts
	if numSeconds >= 60 {
		rate := int64(counthist.apiCounter) / numSeconds
		firstObj := rate * 60
		counthist.ts = counthist.ts + firstObj
		UpdateCountHist(counthist, counterName)
		return true
	}
	log.Printf("Counthist %v", counthist)
	fmt.Printf("Counthist %v", counthist)
	UpdateCountHist(counthist, counterName)
	return false
}

func WriteComputeData(readRange string, values [][]interface{}) {
	sheetId := GetComputeSheetId()

	WriteData(sheetId, readRange, values)
}

func ReadComputeData() {
	sheetId := GetComputeSheetId()
	ReadData(sheetId, DIMENSION_COLUMNS)
}

func WriteResultData(readRange string, values [][]interface{}) {
	sheetId := GetResultSheet()
	WriteData(sheetId, readRange, values)
}

func WriteData(sheetId string, readRange string, values [][]interface{}) {

	client := auth.GetClientInst()
	if client == nil {
		log.Fatalf("Unable to fetch HTTP client")
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}

	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  readRange,
		Values: values,
	})
	_, err = srv.Spreadsheets.Values.BatchUpdate(sheetId, rb).Context(context.Background()).Do()
	if err != nil {
		log.Fatal(err)
	}

	ThrottleApiRequest("WriteData")

}

func ReadData(sheetId string, majorDimension string) *sheets.ValueRange {

	client := auth.GetClientInst()
	if client == nil {
		log.Fatalf("Unable to fetch HTTP client")
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(sheetId, "Sheet1").MajorDimension(majorDimension).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read Data Complete")
	ThrottleApiRequest("ReadData")

	return resp

}

func PushCommand(symbol string, sheetId string) {

	values := make([][]interface{}, 1)
	secondElem := make([]interface{}, 2)
	gFinString := gfin.GetGFinCommand(symbol,
		gfin.PRICE, utils.THREEYEARS, gfin.INTERVAL_DAILY)
	secondElem[0] = gFinString
	values[0] = secondElem
	readRange := "Sheet1!A1:1"
	WriteData(sheetId, readRange, values)

}

type Matrix [][]interface{}

func Push2DArray(array Matrix, sheetId string, sheetName string) {

	readRange := fmt.Sprintf("%s", sheetName)
	WriteData(sheetId, readRange, array)
}

func GetFloat(value interface{}) (float64, error) {
	switch i := value.(type) {
	case float64:
		return float64(i), nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case string:
		s, _ := strconv.ParseFloat(i, 64)
		return s, nil
	default:
		log.Printf("Type %v\n", i)
		return math.NaN(), errors.New("getFloat: unknown value is of incompatible type")
	}

}

func FetchData(sheetId string) []models.ScriptElem {
	//Fetch The result
	rsp := ReadData(sheetId, DIMENSION_ROWS)
	values := rsp.Values
	scriptElems := []models.ScriptElem{}
	for _, item := range values {
		if item[0] == "Date" {
			continue
		}
		if len(item) <= 1 {
			continue
		}
		floatValue, err := GetFloat(item[1])
		if err != nil {
			log.Print("Unable to convert the float value %s", err)
		}
		timeElem, err := time.Parse("1/2/2006 15:04:05", fmt.Sprintf("%s", item[0]))
		scriptElems = append(scriptElems,
			models.ScriptElem{Date: timeElem, Value: floatValue})
	}
	return scriptElems
}

func ProcessScriptInfo(script ScriptInfo) models.ScriptData {
	//Write the command
	sheetId := GetWorkerSheet()
	scriptElems := []models.ScriptElem{}
	fmt.Printf("Pushing Commands ")
	PushCommand(script.Symbol, sheetId)

	readChan := make(chan []models.ScriptElem)

	go func() {
		scriptElems := FetchData(sheetId)
		readChan <- scriptElems
	}()

	select {
	case scriptElems = <-readChan:
	case <-time.After(30 * time.Second):
		fmt.Println("out of time :(")
	}

	fmt.Printf("Returning Commands ")
	return models.ScriptData{
		Symbol:     script.Symbol,
		NumObj:     int64(len(scriptElems)),
		ScriptHist: scriptElems}
}

func GenerateResultSheet() {
	resultHash := utils.GetResultHash()

	log.Printf("GenerateResultJson Size of result Hash %d\n", resultHash.Size())
	var outArray = make([][]interface{}, resultHash.Size())
	outCsvStr := ""
	for key, results := range resultHash.Items {
		fmt.Printf("Processing Key %v len %d\n", key, len(results))
		scriptScoreArray := make([]interface{}, len(results)+2)
		symbol := ""
		for _, resO := range results {
			res := resO.(models.StatResult)
			symbol = res.Symbol
			scriptScoreArray[0] = symbol
			idx := stats.GetIntervalId(res.Name)
			if idx == -1 {
				fmt.Printf("Invalid Algo %s", res.Name)
				continue
			}
			fmt.Printf("Idex %d Name %v\n", idx, res.Name)
			scriptScoreArray[idx] = fmt.Sprintf("%f", res.Score)
		}
		outArray = append(outArray, scriptScoreArray)
	}

	Push2DArray(outArray, GetResultsSheet(), "Score")
	fmt.Printf("%v", outCsvStr)
}
