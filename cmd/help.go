package cmd

import "fmt"

func showHelp() {
	fmt.Println("\t\t\tAVAILABLE COMMMANDS")
	fmt.Println("\thelp \t\t\t\t- Show this help message")
	fmt.Println("\tview [month] \t\t\t- View financial data for the given month")
	fmt.Println("\tadd [month] [day] [income] \t\t\t- Add income to specific day of a month")
	fmt.Println("\texit \t\t\t\t- Quit program")
}
