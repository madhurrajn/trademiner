package score

import (
	"context"
	"fmt"
	"log"
	"miner/auth"
	"miner/excel"
	"miner/gfin"
	"time"

	"google.golang.org/api/sheets/v4"
)

func initSpreadSheet() {

	spreadsheetId := "1h7aG-c_nFF3qCr4wXVkhzU0HOmofSY_Bq0wtpOxyhyM"

	client := auth.GetClientInst()
	if client == nil {
		log.Fatalf("Unable to fetch http client")
	}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	readRange := "Compute!A1:E"

	var market_name = "NSE"
	allScripts := excel.GetAllNseScripts()
	numObject := len(allScripts)
	values := make([][]interface{}, numObject)
	numDays := 7
	for i, v := range allScripts {
		secondElem := make([]interface{}, 1)
		year, month, day := time.Now().AddDate(0, 0, -numDays).Date()
		ago_date := fmt.Sprintf("%d/%d/%d", day, int(month), year)
		full_name := fmt.Sprintf("%s:%s", market_name, v.Symbol)
		secondElem[0] = fmt.Sprintf("=GOOGLEFINANCE(\"%s\",\"price\", \"%s\", %d, \"daily\")", full_name, ago_date, numDays)
		values[i] = secondElem
		rb := &sheets.BatchUpdateValuesRequest{
			ValueInputOption: "USER_ENTERED",
		}

		rb.Data = append(rb.Data, &sheets.ValueRange{
			Range:  readRange,
			Values: values,
		})
		_, err = srv.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Context(context.Background()).Do()
		if err != nil {
			log.Fatal(err)
		}

		break
	}

}

func Refresh() {

	excel.Init()

	allScripts := excel.GetAllNseScripts()

	values := make([][]interface{}, 1)

	secondElem := make([]interface{}, len(allScripts)*2)

	gFinString := ""
	colIdx := 0
	fmt.Printf("Calling Refresh")
	for _, script := range allScripts {
		gFinString = gfin.GetGFinCommand(script.Symbol, gfin.PRICE,
			gfin.THREEYEARS, gfin.INTERVAL_DAILY)
		secondElem[colIdx] = gFinString
		colIdx++
		secondElem[colIdx] = ""
		colIdx++
		//fmt.Printf("%s %d %s", secondElem, colIdx, gFinString)
	}
	//fmt.Printf("%s", secondElem)
	values[0] = secondElem

	readRange := "Sheet1!A1:1"
	excel.WriteComputeData(readRange, values)

}

func Run() {
	excel.Init()
	allScripts := excel.GetAllNseScripts()

	initSpreadSheet()
	for _, script := range allScripts {
		fmt.Printf("%v\n", script)
	}
	fmt.Println("vim-go")
}
