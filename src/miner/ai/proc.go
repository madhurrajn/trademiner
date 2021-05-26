package ai

import (
	"fmt"
	"miner/ai/stats"
	"miner/models"
	"miner/utils"
	"sync"
)

var StatAlgo = [...]int{utils.ONEDAY, utils.ONEWEEK, utils.FIFTEENDAYS,
	utils.ONEMONTH, utils.TWOMONTHS, utils.THREEYEARS,
	utils.SIXMONTH, utils.ONEYEAR, utils.TWOYEARS,
	utils.THREEYEARS}

func ProcDataStream(procData models.ScriptData, resultStream chan<- models.StatResult) float64 {

	var wg sync.WaitGroup
	fmt.Printf("\n Processing Data Stream Symbol %s\n", procData.Symbol)
	for _, statalg := range StatAlgo {
		wg.Add(1)
		scriptElems := procData.ScriptHist
		if len(scriptElems) > 1 {
			go func(statalg int) { defer wg.Done(); resultStream <- stats.GetIntervalScore(procData, statalg) }(statalg)
		} else {
			wg.Done()
		}
	}
	wg.Wait()

	return 0.0

}
