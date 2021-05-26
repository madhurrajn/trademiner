package main

import (
	"log"
	"miner/auth"
	"miner/excel"
	"os"
	"testing"
)

func TestRefresh(t *testing.T) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	auth.Init()

	excel.ReadComputeData()
}
