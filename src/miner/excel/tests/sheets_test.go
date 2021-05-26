package main

import (
	"fmt"
	"miner/auth"
	"miner/excel"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	auth.Init()
	sheetId := excel.GetComputeSheetId()

	rangeData := fmt.Sprintf("%s!$A1", "Sheet1")

	values := make([][]interface{}, 1)
	inner := make([]interface{}, 1)
	inner[0] = fmt.Sprintf("HelloWorld")
	values[0] = inner

	excel.WriteData(sheetId, rangeData, values)
}
