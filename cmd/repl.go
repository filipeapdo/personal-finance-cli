package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartREPL() {
	fmt.Println("Welcome to the Personal Finance CLI!")
	fmt.Println("Try 'help' to see available commands or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pfc> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("Exiting... Goodbye!")
			break
		}

		handleCommand(input)
	}
}

func handleCommand(input string) {
	parts := strings.Fields(input)
	command := parts[0]

	switch command {
	case "help":
		showHelp()
	default:
		fmt.Println("Unknown command.")
		fmt.Println("Try 'help' to see available commands or 'exit' to quit.")
	}
}

func showHelp() {
	fmt.Println("\t\t\tAVAILABLE COMMMANDS")
	fmt.Println("\thelp \t\t\t\t- Show this help message")
	fmt.Println("\tview [month] \t\t\t- View all data for a specific month")
	fmt.Println("\tset income [day] [amount] \t- Set planned income for a day")
	fmt.Println("\tset expense [day] [amount] \t- Set planned expense for a day")
	fmt.Println("\tset daily [amount] \t\t- Set fixed daily expense for future days")
	fmt.Println("\tsummary [month] \t\t- Show summary for a specific month")
	fmt.Println("\texit \t\t\t\t- Quit program")
}
