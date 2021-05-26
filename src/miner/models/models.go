package models

import (
	"time"
)

type ScriptElem struct {
	Date  time.Time
	Value float64
}

type ScriptData struct {
	Symbol     string
	NumObj     int64
	ScriptHist []ScriptElem
}

const (
	HIGH = iota
	LOW
	MEDIUM
	BAILOUT
)

type StatResult struct {
	Symbol     string  `json:symbol`
	Name       string  `json:name`
	High       float64 `json:high`
	Low        float64 `json:low`
	Avg        float64 `json:avg`
	Score      float64 `json:score`
	Confidence int     `json:confidence`
}

type CompositeResult struct {
	Symbol     string       `json:symbol`
	StatResult []StatResult `json:statresult`
}
