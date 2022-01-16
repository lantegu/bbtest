package datascript

import (
	"encoding/csv"
	"errors"
	"os"
)

func ReadCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Unable to read input file "+filePath)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("Unable to parse file as CSV for "+filePath)
	}
	return records, nil
}