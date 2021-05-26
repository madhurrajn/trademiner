package sa

import (
	"miner/ai"
	"miner/ai/stats"
	"miner/db"
	"miner/excel"
	"miner/models"
	"sync"
)

func getResultsInfo(symbol string, intName string) models.StatResult {
	obj := db.ReadResultsInfo(symbol, intName)
	if len(obj) > 0 {
		return obj[0]
	} else {
		return models.StatResult{}
	}
}

func ProcessScriptInfo(scriptInfo excel.ScriptInfo, resultStream chan<- models.StatResult) float64 {

	var wg sync.WaitGroup
	for _, statAlgo := range ai.StatAlgo {
		wg.Add(1)
		go func(statAlgo int) {
			defer wg.Done()
			resultStream <- getResultsInfo(scriptInfo.Symbol, stats.GetIntervalName(statAlgo))
		}(statAlgo)
	}
	wg.Wait()
	return 0.0
}
