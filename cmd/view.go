package cmd

import (
	"fmt"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func viewMonth(month string) {
	m, err := findMonthByName(month)
	if err != nil {
		fmt.Println(err)
		return
	}

	printMonthSummary(m)
}

func findMonthByName(month string) (*data.Month, error) {
	for _, m := range financeData.Months {
		if m.Name == month {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("no data available for month %s", month)
}

func printMonthSummary(month *data.Month) {
	fmt.Printf("\nFinancial data for %s:\n", month.Name)
	fmt.Println("Day | Income  | Expense | Daily   | Balance")
	fmt.Println("-----------------------------------------")
	for _, day := range month.Days {
		fmt.Printf("%2d  | %7.2f | %7.2f | %7.2f | %7.2f\n",
			day.Day, day.Income, day.Expense, day.Daily, day.Balance)
	}
	fmt.Println("\nSummary:")
	fmt.Printf("Total Income:  %.2f\n", month.Summary.TotalIncome)
	fmt.Printf("Total Expense: %.2f\n", month.Summary.TotalExpense)
	fmt.Printf("Total Daily:   %.2f\n", month.Summary.TotalDaily)
	fmt.Printf("Final Balance: %.2f\n\n", month.Summary.FinalBalance)
}
