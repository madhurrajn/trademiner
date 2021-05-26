package main

import (
	"encoding/json"
	"fmt"
	"log"
	"miner/db"
	"testing"
)

type TestObjectSub2 struct {
	SubInt2    int    `json:"subint2"`
	SubString2 string `json:"substring2"`
}

type TestObject struct {
	SubInt1        int            `json:"subint1"`
	SubString1     string         `json:"subint2"`
	TestObjectSub2 TestObjectSub2 `json:"testobjectsub2"`
}
type TestObjectArray struct {
	SubInt1        int              `json:"subint1"`
	SubString1     string           `json:"subint2"`
	TestObjectSub2 []TestObjectSub2 `json:"testobjectsub2"`
}

func TestScribbleDb(t *testing.T) {
	db.Init("scribble")

	dbObj := TestObjectArray{
		SubInt1:    1,
		SubString1: "string1",
		TestObjectSub2: []TestObjectSub2{TestObjectSub2{
			SubInt2:    2,
			SubString2: "string2",
		},
			TestObjectSub2{
				SubInt2:    3,
				SubString2: "string2",
			},
		},
	}

	db.Write("testobject", "key1", dbObj)

	if obj, err := json.Marshal(dbObj); err != nil {
		fmt.Printf("Unable to print obj")
	} else {
		fmt.Printf("%s ", obj)
	}
}

func TestScribbleDbArray(t *testing.T) {
	db.Init("scribble")

	dbObj := TestObject{
		SubInt1:    1,
		SubString1: "string1",
		TestObjectSub2: TestObjectSub2{
			SubInt2:    2,
			SubString2: "string2",
		},
	}

	db.Write("testobject", "key2", dbObj)

	if obj, err := json.Marshal(dbObj); err != nil {
		fmt.Printf("Unable to print obj")
	} else {
		fmt.Printf("%s ", obj)
	}
}

func TestScribbleRead(t *testing.T) {
	db.Init("scribble")

	dbObj := TestObject{
		SubInt1:    1,
		SubString1: "string1",
		TestObjectSub2: TestObjectSub2{
			SubInt2:    2,
			SubString2: "string2",
		},
	}

	db.Write("testobject", "key2", dbObj)

	for _, obj := range db.Read("NSE:TCS", "ONEYEAR") {
		fmt.Println(obj)
	}
}

func TestScribbleReadDbArray(t *testing.T) {
	db.Init("scribble")

	dbObj := TestObjectArray{
		SubInt1:    1,
		SubString1: "string1",
		TestObjectSub2: []TestObjectSub2{TestObjectSub2{
			SubInt2:    2,
			SubString2: "string2",
		},
			TestObjectSub2{
				SubInt2:    3,
				SubString2: "string2",
			},
		},
	}

	db.Write("testobject", "key1", dbObj)

	for _, obj := range db.Read("testobject", "key1") {
		log.Printf("%v\n", obj)
	}

}

func TestScribbleReadAll(t *testing.T) {
	db.Init("scribble")
	for _, obj := range db.ReadAll("NSE:TCS") {
		log.Printf("%v\n", obj)
	}

}
