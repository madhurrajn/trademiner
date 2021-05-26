package excel

const (
	COMPUTESHEET     = "1jI2T8FO5Kt7AkUvE6cS3I6Fm8bt2BeuPbXf8UXVBdEk"
	RESULTSHEET      = "1h7aG-c_nFF3qCr4wXVkhzU0HOmofSY_Bq0wtpOxyhyM"
	WORKERSHEET      = "1l0KW9vebKjydaZTI9T04jnCq-bePJ4MaDslgVkelbmg"
	RESULTSSHEET     = "1kf_2FEF39tgJ3zjMDS5Sgf_vUfJDDQr6f-WlKYvxAAM"
	NUMCOMPUTESHEETS = 50
)

const (
	DIMENSION_COLUMNS = "COLUMNS"
	DIMENSION_ROWS    = "ROWS"
)

func GetComputeSheetId() string {
	return COMPUTESHEET
}

func GetResultSheet() string {
	return RESULTSHEET
}

func GetNumComputeSheet() int {
	return NUMCOMPUTESHEETS
}

func GetWorkerSheet() string {
	return WORKERSHEET
}

func GetResultsSheet() string {
	return RESULTSSHEET
}
