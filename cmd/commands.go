package cmd

import (
	"errors"
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
	if len(parts) < 4 {
		return errors.New("Usage: add [month] [day] [income]")
	}

	month := parts[1]
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid day: %v", err)
	}

	income, err := strconv.ParseFloat(parts[3], 64)
	if err != nil || income < 0 {
		return fmt.Errorf("invalid income: %v", err)
	}

	err = AddIncome(financeData, month, day, income)

	fmt.Printf("Successfully added %.2f income to %s, day %d.\n", income, month, day)
	return nil
}
