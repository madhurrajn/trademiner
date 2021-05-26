package results

import (
	"encoding/json"
	"fmt"
	"log"
	"miner/ai/stats"
	"miner/models"
	"miner/utils"
	"os"
	"strings"
)

func ProcResultsInfo(resultInfo models.StatResult) float64 {
	resultHash := utils.GetResultHash()

	log.Printf("Recevied Stat Result %v", resultInfo)
	resultHash.Put(resultInfo.Symbol, resultInfo)

	log.Printf("Size of result Hash %d\n", resultHash.Size())

	return 0.0
}

func GenerateResultJson() {
	resultHash := utils.GetResultHash()

	log.Printf("GenerateResultJson Size of result Hash %d\n", resultHash.Size())
	outStrJson := ""
	for _, results := range resultHash.Items {
		for _, res := range results {
			outStr, err := json.Marshal(res)
			if err != nil {
				log.Println("Failed to Marshal Json %v", err)
			}
			if outStrJson == "" {
				outStrJson = fmt.Sprintf("[%s", outStr)
			} else {
				outStrJson = fmt.Sprintf("%s,%s", outStrJson, outStr)
			}
		}
		outStrJson = fmt.Sprintf("%s]", outStrJson)
	}
	log.Println(outStrJson)
}

func GenerateResultCsv() {
	resultHash := utils.GetResultHash()

	log.Printf("GenerateResultJson Size of result Hash %d\n", resultHash.Size())
	var finArray [][]float64
	outCsvStr := ""
	for key, results := range resultHash.Items {
		fmt.Printf("Processing Key %v\n", key)
		scriptScoreArray := [utils.MAX_STAT_ALGO + 1]string{}
		symbol := ""
		for _, resO := range results {
			res := resO.(models.StatResult)
			symbol = res.Symbol
			idx := stats.GetIntervalId(res.Name)
			if idx == -1 {
				fmt.Printf("Invalid Algo %s", res.Name)
				continue
			}
			scriptScoreArray[idx] = fmt.Sprintf("%f", res.Score)
		}
		outCsvStr = outCsvStr + symbol
		outCsvStr = outCsvStr + strings.Join(scriptScoreArray[:], ",")
		outCsvStr = outCsvStr + "\n"
	}

	f, _ := os.Create("users.csv")
	defer f.Close()
	f.WriteString(outCsvStr)

	fmt.Printf("%v", outCsvStr)
	log.Println(finArray)
}
