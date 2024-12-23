package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func StartREPL(fd *data.FinanceData) {
	fmt.Println("Welcome to the Personal Finance CLI!")
	fmt.Println("Try 'help' to see available commands or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input: ", err)
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

		handleCommand(fd, input)
	}
}
