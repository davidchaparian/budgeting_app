package budgeting

import (
	"encoding/csv"
	"os"
)

func OpenCSVFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadCSV(file *os.File) ([][]string, error) {
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func ProcessCSVData(records [][]string) {
	for _, record := range records {
		for _, value := range record {
			print(value)
		}
	}
}
