package gfin

import "fmt"

const (
	PRICE     = "price"
	PRICEOPEN = "priceopen"
	HIGH      = "high"
	LOW       = "low"
	VOLUME    = "volume"
	MARKETCAP = "marketcap"
	TRADETIME = "tradetime"
	DATADELAY = "datadelay"
	VOLUMEAVG = "volumeavg"
	PE        = "pe"
	EPS       = "eps"
	HIGH52    = "high52"
	LOW52     = "low52"
	CHANGE    = "change"
	BETA      = "beta"
	CHANGEPCT = "changepct"
	CLOSEYEST = "closeyest"
	SHARES    = "shares"
	CURRENCY  = "currency"
)

const (
	INTERVAL_DAILY  = "daily"
	INTERVAL_WEEKLY = "weekly"
)

func GetGFinCommand(symbol string, attrib string, numdays int, interval string) string {

	start_date := fmt.Sprintf("today()-%d", numdays)
	end_date := fmt.Sprintf("today()")
	gfinString := ""
	if numdays == 0 {
		gfinString = fmt.Sprintf("=GOOGLEFINANCE(\"%s\", \"%s\")",
			symbol, attrib)
	} else {
		gfinString = fmt.Sprintf("=GOOGLEFINANCE(\"%s\", \"%s\", %s, %s, \"%s\")",
			symbol, attrib, start_date, end_date, interval)
	}

	return gfinString
}
