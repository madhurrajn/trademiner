package db

import (
	"encoding/json"
	"fmt"
	"log"
	"miner/db/scribledb"
	"miner/models"

	scribble "github.com/nanobox-io/golang-scribble"
)

var dbContext interface{}
var dbTypeElem string

func InitConnection(dbType string) (interface{}, error) {

	switch dbType {
	case "scribble":
		if dbElem, err := scribledb.InitScribbleDb(); err != nil {
			fmt.Println("Unable to Init Scribble db")
		} else {
			dbContext = dbElem
			dbTypeElem = dbType
			return dbElem, err
		}
	}
	return nil, fmt.Errorf("Unable to select DB")

}

func Write(conn interface{}, key interface{}, object interface{}) {
	switch dbContext.(type) {
	case *scribble.Driver:
		if err := scribledb.Write(conn, key, object); err != nil {
			fmt.Println("Error ", err)
		}
	}
}

func WriteResultsInfo(resultStat models.StatResult) float64 {
	key := fmt.Sprintf("%s", resultStat.Name)
	resource := fmt.Sprintf("%s", resultStat.Symbol)
	Write(resource, key, resultStat)
	return 0.0
}

func Read(resource string, key string) []models.StatResult {
	object := models.StatResult{}
	switch dbContext.(type) {
	case *scribble.Driver:
		newObj := models.StatResult{}
		if obj, err := scribledb.Read(resource, key, &newObj); err != nil {
			fmt.Println("Error ", err)
		} else {
			if obj != nil {
				object = newObj
			}
		}
	}
	res := []models.StatResult{}
	res = append(res, object)
	return res
}

func ReadAll(resource string) []models.StatResult {
	res := []models.StatResult{}
	switch dbContext.(type) {
	case *scribble.Driver:
		if records, err := scribledb.ReadAll(resource); err != nil {
			fmt.Println("Error ", err)
		} else {
			for _, record := range records {
				statResult := models.StatResult{}
				if err := json.Unmarshal([]byte(record), &statResult); err != nil {
					fmt.Println("Error", err)
				}
				res = append(res, statResult)
			}
		}
	}
	return res
}

func ReadResultsInfo(resource string, key string) []models.StatResult {

	if key != "" {
		return Read(resource, key)
	} else {
		return ReadAll(resource)
	}

}

func Init(dbType string) {

	if db, err := InitConnection(dbType); err != nil {
		log.Println("Unable Initialize Db ", db)
	}
}
