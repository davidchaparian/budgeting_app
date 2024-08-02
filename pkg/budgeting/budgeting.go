package budgeting

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

func MostExpense(data [][]string) {
	mostExpenseStr := data[1][2]
	indexOfMostExpense := 0
	mostExpense, err := strconv.Atoi(mostExpenseStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	for j := 2; j < 10; j++ {
		currentExpenseStr := data[j][2]
		currentExpense, err := strconv.Atoi(currentExpenseStr)
		if err != nil {
			fmt.Println("Error converting string to float: ", err)
			continue
		}
		if currentExpense > mostExpense {
			mostExpense = currentExpense
			indexOfMostExpense = j
		}

	}
	nameOfMostExpense := data[indexOfMostExpense][0]
	fmt.Println("The most expense: ", nameOfMostExpense, "=", mostExpense)
}
