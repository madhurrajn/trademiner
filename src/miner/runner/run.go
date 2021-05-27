package runner

import (
	"fmt"
	"log"
	"miner/ai/sa"
	"miner/auth"
	"miner/db"
	"miner/excel"
	"miner/results"
	"miner/utils"
	"os"
)

func Run() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	excel.Init()
	utils.Init()
	auth.Init()
	db.Init("scribble")
	log.Println("Auth Initialized")
	fmt.Printf("Auth Initialized")

	allNseScripts := excel.GetAllNseScripts()
	allNseScripts = allNseScripts[:2]
	done := make(chan interface{})
	defer close(done)
	scriptInfoStream := excel.ScriptGenerator(done, allNseScripts)
	log.Printf("Script Generator %v", scriptInfoStream)

	dataStream := excel.GetData(done, scriptInfoStream)

	resultStream := excel.ProcDataStream(done, dataStream)

	dbStream := excel.DbStream(done, resultStream)

	for ov := range dbStream {
		log.Printf("%v\n", ov)
	}

	results.GenerateResultJson()
}

func RunAiStats() {
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
	done := make(chan interface{})
	defer close(done)
	scriptInfoStream := excel.ScriptGenerator(done, allNseScripts)
	log.Printf("Script Generator %v", scriptInfoStream)

	resultStream := sa.ReadDbData(done, scriptInfoStream)

	doneStream := excel.ResultStatStream(done, resultStream)

	for ov := range doneStream {
		log.Printf("%v\n", ov)
	}

	excel.GenerateResultSheet()
	results.GenerateResultCsv()
}
