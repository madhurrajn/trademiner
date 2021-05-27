package excel

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ScriptInfo struct {
	Symbol  string
	Company string
	Market  int64
}

var allScripts []ScriptInfo

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func GetAllNseScripts() []ScriptInfo {
	return allScripts
}

func Init() {
	f, err := excelize.OpenFile("config/nse.xlsx")
	if err != nil {
		f, err = excelize.OpenFile("config/nse.xlsx")
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("NSE")
	skip := false
	for _, row := range rows {
		script_data := ScriptInfo{}
		skip = false
		for idx, colCell := range row {
			switch idx {
			case 1:
				script_data.Symbol = fmt.Sprintf("NSE:%s", colCell)
			case 2:
				script_data.Company = colCell
			case 3:
				if isInt(colCell) {
					script_data.Market, err = strconv.ParseInt(colCell, 10, 64)
				} else {
					skip = true
					break
				}

			default:

			}
		}
		if skip == false {
			allScripts = append(allScripts, script_data)
		}
	}

}

func main() {
	f, err := excelize.OpenFile("nse.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("NSE", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("NSE")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	Init()
}
