package cmd

import (
	"fmt"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func insertAmount(fd *data.FinanceData, insertType, monthName string, day int, amount float64) error {
	month, err := findOrCreateMonth(fd, monthName)
	if err != nil {
		return err
	}

	dayIndex, err := findOrCreateDay(month, day)
	if err != nil {
		return err
	}

	switch insertType {
	case "income":
		month.Days[dayIndex].Income += amount
	case "expense":
		month.Days[dayIndex].Expense += amount
	case "daily":
		month.Days[dayIndex].Daily += amount
	default:
		return fmt.Errorf("invalid type: must be 'income', 'expense' or 'daily'")
	}
	fmt.Printf("Successfully inserted %.2f %s to %s, day %d.\n", amount, insertType, monthName, day)

	// to-do: review with error handlings
	data.SortFinanceData(fd)
	calculateBalance(fd)
	calculateSummary(fd)

	err = data.SaveFinanceData(fd)
	if err != nil {
		return err
	}

	return nil
}

func findOrCreateMonth(fd *data.FinanceData, monthName string) (*data.Month, error) {
	err := data.ValidateMonth(monthName)
	if err != nil {
		return &data.Month{}, err
	}

	for i := range fd.Months {
		if fd.Months[i].Name == monthName {
			return &fd.Months[i], nil
		}
	}

	newMonth := data.Month{
		Name: monthName,
		Days: []data.Day{},
		Summary: struct {
			TotalIncome  float64 `json:"total_income"`
			TotalExpense float64 `json:"total_expense"`
			TotalDaily   float64 `json:"total_daily"`
			FinalBalance float64 `json:"final_balance"`
		}{},
	}
	fd.Months = append(fd.Months, newMonth)
	return &fd.Months[len(fd.Months)-1], nil
}

func findOrCreateDay(month *data.Month, day int) (int, error) {
	// to-do: review this function!!!
	err := data.ValidateDay(month.Name, day, 2024)
	if err != nil {
		return -1, err
	}

	for i := range month.Days {
		if month.Days[i].Day == day {
			return i, nil
		}
	}

	newDay := data.Day{
		Day:     day,
		Income:  0,
		Expense: 0,
		Daily:   0,
		Balance: 0,
	}
	month.Days = append(month.Days, newDay)

	return len(month.Days) - 1, nil
}
