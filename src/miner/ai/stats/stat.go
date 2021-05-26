package stats

import (
	"log"
	"miner/models"
	"miner/utils"
	"time"
)

func getDateFromInterval(interval int) time.Time {
	return time.Now().AddDate(0, 0, -1*interval)
}

func getScriptSlice(scriptElems []models.ScriptElem, startDate time.Time) []models.ScriptElem {
	exitIdx := 0
	objDate := time.Now()
	for idx, obj := range scriptElems {
		objDate = obj.Date
		if objDate.After(startDate) {
			exitIdx = idx
			break
		}
	}
	log.Printf("objDate %v startdate %v Sliced Index %d \n", objDate, startDate, exitIdx)
	return scriptElems[exitIdx:]
}

func getMinMaxAvg(scriptElem []models.ScriptElem) (float64, float64, float64) {
	var min float64 = scriptElem[0].Value
	var max float64 = scriptElem[0].Value
	var avg float64 = scriptElem[0].Value
	var sum float64 = 0
	for _, elem := range scriptElem {
		if min > elem.Value {
			min = elem.Value
		}
		if max < elem.Value {
			max = elem.Value
		}
		sum = sum + elem.Value
	}
	avg = sum / float64(len(scriptElem))
	return min, max, avg
}

func getPercentHike(firstPrice float64, diff float64) float64 {
	return 100 * diff / firstPrice
}

func getStatScore(scriptElem []models.ScriptElem) models.StatResult {
	min, max, avg := getMinMaxAvg(scriptElem)
	log.Printf("Min %v Max %v Avg %v\n", min, max, avg)
	score := float64(0)
	counter := 0

	if max-min > avg {
		score += 10
		counter = counter + 1
	}

	arrayLen := len(scriptElem)
	firstPrice := scriptElem[0].Value
	lastPrice := scriptElem[arrayLen-1].Value

	if lastPrice > firstPrice {
		score = score + 10
		counter = counter + 1
	}

	hike := getPercentHike(firstPrice, (lastPrice - firstPrice))
	hike10thValue := hike / 10
	score = score + hike10thValue
	counter = counter + 1

	log.Printf("firstPrice %v lastPrice %v hike %v hike10th %v\n", firstPrice, lastPrice, hike, hike10thValue)
	finScore := score / float64(counter)

	return models.StatResult{
		High:       max,
		Low:        min,
		Avg:        avg,
		Score:      finScore,
		Confidence: models.HIGH,
	}

}

func GetIntervalId(interval string) int {
	switch interval {
	case "ONEDAY":
		return utils.ONEDAYIDX
	case "ONEWEEK":
		return utils.ONEWEEKIDX
	case "FIFTEENDAYS":
		return utils.FIFTEENDAYSIDX
	case "ONEMONTH":
		return utils.ONEMONTHIDX
	case "TWOMONTHS":
		return utils.TWOMONTHSIDX
	case "THREEMONTHS":
		return utils.THREEMONTHSIDX
	case "SIXMONTH":
		return utils.SIXMONTHIDX
	case "ONEYEAR":
		return utils.ONEYEARIDX
	case "TWOYEARS":
		return utils.TWOYEARSIDX
	case "THREEYEARS":
		return utils.THREEYEARSIDX
	}
	return -1
}

func GetIntervalName(interval int) string {
	switch interval {
	case utils.ONEDAY:
		return "ONEDAY"
	case utils.ONEWEEK:
		return "ONEWEEK"
	case utils.FIFTEENDAYS:
		return "FIFTEENDAYS"
	case utils.ONEMONTH:
		return "ONEMONTH"
	case utils.TWOMONTHS:
		return "TWOMONTHS"
	case utils.THREEMONTHS:
		return "THREEMONTHS"
	case utils.SIXMONTH:
		return "SIXMONTH"
	case utils.ONEYEAR:
		return "ONEYEAR"
	case utils.TWOYEARS:
		return "TWOYEARS"
	case utils.THREEYEARS:
		return "THREEYEARS"
	}
	return ""
}

func GetIntervalScore(procData models.ScriptData, interval int) models.StatResult {
	startDate := getDateFromInterval(interval)
	//endDate := time.Now()

	scriptElems := procData.ScriptHist

	slice := getScriptSlice(scriptElems, startDate)

	statResult := getStatScore(slice)
	statResult.Symbol = procData.Symbol
	statResult.Name = GetIntervalName(interval)
	return (statResult)

}
