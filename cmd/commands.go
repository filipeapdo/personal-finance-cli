package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func handleCommand(input string) {
	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "help":
		showHelp()
	case "view":
		handleView(parts)
	case "add":
		err := handleAddCommand(parts, &financeData)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Unknown command.")
		fmt.Println("Try 'help' to see available commands or 'exit' to quit.")
	}
}

func handleView(parts []string) {
	if len(parts) < 2 {
		fmt.Println("Usage: view [month].")
		return
	}
	month := parts[1]
	viewMonth(month)
}

func handleAddCommand(parts []string, financeData *data.FinanceData) error {
	if len(parts) < 5 {
		return fmt.Errorf("Usage: add [income | expense] [month] [day] [amount]")
	}

	addType := parts[1]
	month := parts[2]
	day, err := strconv.Atoi(parts[3])
	if err != nil {
		return fmt.Errorf("invalid day: %v", err)
	}

	amount, err := strconv.ParseFloat(parts[4], 64)
	if err != nil || amount < 0 {
		return fmt.Errorf("invalid income: %v", err)
	}

	switch addType {
	case "income":
		err = AddIncome(financeData, month, day, amount)
		if err != nil {
			return fmt.Errorf("error adding income: %v", err)
		}
		fmt.Printf("Successfully added %.2f income to %s, day %d.\n", amount, month, day)
	case "expense":
		err = AddExpense(financeData, month, day, amount)
		if err != nil {
			return fmt.Errorf("error adding expense: %v", err)
		}
		fmt.Printf("Successfully added %.2f expense to %s, day %d.\n", amount, month, day)
	default:
		return fmt.Errorf("invalid type: must be 'income' or 'expense'")
	}

	return nil
}
