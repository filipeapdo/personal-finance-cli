package cmd

import (
	"fmt"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func addAmount(fd *data.FinanceData, addType, monthName string, day int, amount float64) error {
	month, err := findOrCreateMonth(fd, monthName)
	if err != nil {
		return err
	}

	dayIndex, err := findOrCreateDay(month, day)
	if err != nil {
		return err
	}

	switch addType {
	case "income":
		month.Days[dayIndex].Income += amount
		fmt.Printf("Successfully added %.2f income to %s, day %d.\n", amount, monthName, day)
	case "expense":
		month.Days[dayIndex].Expense += amount
		fmt.Printf("Successfully added %.2f expense to %s, day %d.\n", amount, monthName, day)
	case "daily":
		month.Days[dayIndex].Daily += amount
		fmt.Printf("Successfully added %.2f daily expense to %s, day %d.\n", amount, monthName, day)
	default:
		return fmt.Errorf("invalid type: must be 'income', 'expense' or 'daily'")
	}

	err = data.SaveFinanceData(fd)
	if err != nil {
		return err
	}

	return nil
}

func findOrCreateMonth(fd *data.FinanceData, monthName string) (*data.Month, error) {
	err := validateMonth(monthName)
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

func validateMonth(monthName string) error {
	monthNames := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	for _, m := range monthNames {
		if m == monthName {
			return nil
		}
	}

	return fmt.Errorf("invalid month name, valid month's names are: %v", monthNames)
}

func findOrCreateDay(month *data.Month, day int) (int, error) {
	err := validateDay(day)
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

// to-do: validate the day based on months, 31/30/28or29
func validateDay(day int) error {
	if day < 1 || day > 31 {
		return fmt.Errorf("invalid day, valid days are: between 1 and 31")
	}

	return nil
}
