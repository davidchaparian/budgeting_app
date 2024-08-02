package main

import (
	"fmt"

	"github.com/davidchaparian/budgeting_app/pkg/budgeting"
)

func main() {
	file, err := budgeting.OpenCSVFile("../budget.csv")
	if err != nil {
		fmt.Println("Error opening CSV file: ", err.Error())
	}
	records, err := budgeting.ReadCSV(file)
	if err != nil {
		fmt.Println("Error reading CSV file: ", err.Error())
		return
	}

	budgeting.MostExpense(records)
}
