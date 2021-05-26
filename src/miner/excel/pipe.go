package excel

import (
	"fmt"
	"log"
	"miner/ai"
	"miner/db"
	"miner/models"
	"miner/results"
)

var ScriptGenerator = func(done <-chan interface{}, allNseScripts []ScriptInfo) <-chan ScriptInfo {
	scriptStream := make(chan ScriptInfo)
	go func() {
		defer close(scriptStream)
		for _, script := range allNseScripts {
			select {
			case <-done:
				log.Printf("Received Done")
				return
			case scriptStream <- script:
				fmt.Printf("Processing %s\n", script)
				//time.Sleep(500 * time.Millisecond)
			}
		}
		log.Printf("Exiting AllNSE Scripts %d", len(allNseScripts))
	}()
	return scriptStream
}

var GetData = func(done <-chan interface{},
	scriptStream <-chan ScriptInfo) <-chan models.ScriptData {
	getDataStream := make(chan models.ScriptData)
	go func() {
		defer close(getDataStream)
		for i := range scriptStream {
			select {
			case <-done:
				return
			case getDataStream <- ProcessScriptInfo(i):
			}
		}
	}()
	return getDataStream
}

var ProcDataStream = func(done <-chan interface{},
	dataStream <-chan models.ScriptData) <-chan models.StatResult {
	scoreStream := make(chan models.StatResult)
	go func() {
		defer close(scoreStream)
		for i := range dataStream {
			select {
			case <-done:
				return
			default:
				ai.ProcDataStream(i, scoreStream)
			}
		}
		log.Printf("Exiting go Func")
	}()
	return scoreStream
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

var DbStream = func(done <-chan interface{},
	dbStream <-chan models.StatResult) <-chan float64 {
	doneStream := make(chan float64)
	go func() {
		defer close(doneStream)
		for result := range dbStream {
			select {
			case <-done:
				return
			case doneStream <- db.WriteResultsInfo(result):
			}
		}
	}()
	return doneStream

}
