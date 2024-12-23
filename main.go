package main

import (
	"fmt"
	"log"

	"github.com/filipeapdo/personal-finance-cli/cmd"
	"github.com/filipeapdo/personal-finance-cli/data"
)

// to-do: re-think the data storage module
const dataFilePath = "finance_data.json"

func main() {
	financeData, err := data.LoadFinanceData(dataFilePath)
	if err != nil {
		log.Fatalf("error loading finance data: %v", err)
	}
	fmt.Println("Finance data loaded successfully.")

	cmd.StartREPL(financeData)
}
