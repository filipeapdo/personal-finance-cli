package cmd

import "fmt"

func showHelp() {
	fmt.Println("\t\t\t\t\tAVAILABLE COMMMANDS")
	fmt.Println("\thelp \t\t\t\t\t- Show this help message")
	fmt.Println("\tview <month> \t\t\t\t- View financial data for the given month")
	fmt.Println("\tinsert <income | expense | daily> \t- Insert income or expense to specific day of a month \n\t\t<month> <day> <amount>")
	fmt.Println("\texit \t\t\t\t\t- Quit program")
}
