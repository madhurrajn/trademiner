package sa

import (
	"miner/excel"
	"miner/models"
	"miner/results"
)

var ReadDbData = func(done <-chan interface{},
	scriptStream <-chan excel.ScriptInfo) <-chan models.StatResult {
	readDbStream := make(chan models.StatResult)
	go func() {
		defer close(readDbStream)
		for i := range scriptStream {
			select {
			case <-done:
				return
			default:
				ProcessScriptInfo(i, readDbStream)
			}
		}
	}()
	return readDbStream
}

var ResultStatStream = func(done <-chan interface{},
	resultStream <-chan models.StatResult) <-chan float64 {
	doneStream := make(chan float64)
	go func() {
		defer close(doneStream)
		for result := range resultStream {
			select {
			case <-done:
				return
			case doneStream <- results.ProcResultsInfo(result):
			}
		}
	}()
	return doneStream
}
