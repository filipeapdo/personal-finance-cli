package cmd

import (
	"fmt"
	"strings"
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
		fmt.Println("Command not implemented yet.")
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
