package cmd

import "fmt"

func showHelp() {
	fmt.Println("\t\t\t\t\tAVAILABLE COMMMANDS")
	fmt.Println("\thelp \t\t\t\t\t\t- Show this help message")
	fmt.Println("\tview [month] \t\t\t\t\t- View financial data for the given month")
	fmt.Println("\tadd [income | expense] [month] [day] [amount] \t- Add income or expense to specific day of a month")
	fmt.Println("\texit \t\t\t\t\t\t- Quit program")
}
