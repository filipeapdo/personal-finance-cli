package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func handleCommand(fd *data.FinanceData, input string) {
	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "help":
		showHelp()
	case "view":
		err := handleView(fd, parts)
		if err != nil {
			fmt.Println(err)
		}
	case "insert":
		err := handleInsertCommand(fd, parts)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Unknown command.")
		fmt.Println("Try 'help' to see available commands or 'exit' to quit.")
	}
}

func handleView(fd *data.FinanceData, parts []string) error {
	if len(parts) < 2 {
		return fmt.Errorf("usage: view [month]")
	}
	month := parts[1]
	viewMonth(month, fd)

	return nil
}

func handleInsertCommand(fd *data.FinanceData, parts []string) error {
	if len(parts) < 5 {
		return fmt.Errorf("usage: insert [income | expense | daily] [month] [day] [amount]")
	}

	insertType := parts[1]
	monthName := parts[2]
	day, err := strconv.Atoi(parts[3])
	if err != nil {
		return fmt.Errorf("invalid day: %v", err)
	}

	amount, err := strconv.ParseFloat(parts[4], 64)
	if err != nil || amount < 0 {
		return fmt.Errorf("invalid amount: %v", err)
	}

	err = insertAmount(fd, insertType, monthName, day, amount)
	if err != nil {
		return err
	}

	return nil
}
