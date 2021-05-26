package main

import (
	"log"
	"miner/auth"
	"miner/db"
	"miner/excel"
	"miner/results"
	"miner/utils"
	"os"
	"testing"
)

func TestAiStats(t *testing.T) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Test Started")

	excel.Init()
	utils.Init()
	auth.Init()
	db.Init("scribble")
	log.Println("Auth Initialized")

	allNseScripts := excel.GetAllNseScripts()
	//allNseScripts = allNseScripts[:2]
	//log.Printf("All NSE Scripts %d\n", len(allNseScripts))
	done := make(chan interface{})
	defer close(done)
	scriptInfoStream := excel.ScriptGenerator(done, allNseScripts)
	log.Printf("Script Generator %v", scriptInfoStream)

	dataStream := excel.GetData(done, scriptInfoStream)

	resultStream := excel.ProcDataStream(done, dataStream)

	//doneStream := excel.ResultStatStream(done, resultStream)

	dbStream := excel.DbStream(done, resultStream)

	for ov := range dbStream {
		log.Printf("%v\n", ov)
	}

	//for ov := range doneStream {
	//	log.Printf("%v\n", ov)
	//}

	results.GenerateResultJson()
}
