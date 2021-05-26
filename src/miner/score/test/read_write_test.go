package main

import (
	"log"
	"miner/auth"
	"miner/excel"
	"miner/gfin"
	"miner/utils"
	"os"
	"testing"
)

func TestRefresh(t *testing.T) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	auth.Init()

	//Write Test Column
	sheetId := excel.GetWorkerSheet()
	symbol := "NSE:TCS"
	values := make([][]interface{}, 1)
	secondElem := make([]interface{}, 2)
	gFinString := gfin.GetGFinCommand(symbol, gfin.PRICE, utils.ONEWEEK, gfin.INTERVAL_DAILY)
	secondElem[0] = gFinString
	values[0] = secondElem
	readRange := "Sheet1!A1:1"
	excel.WriteData(sheetId, readRange, values)

	//Fetch The result
	rsp := excel.ReadData(sheetId, excel.DIMENSION_ROWS)
	values = rsp.Values
	for _, item := range values {
		for _, elem := range item {
			log.Println(elem)
		}
	}

}
