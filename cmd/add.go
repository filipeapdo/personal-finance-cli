package cmd

import (
	"fmt"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func AddIncome(financeData *data.FinanceData, monthName string, day int, amount float64) error {
	month, err := findMonth(financeData, monthName)
	if err != nil {
		return err
	}

	if err := validateDay(day, month); err != nil {
		return err
	}

	month.Days[day-1].Income += amount

	return nil
}

func findMonth(financeData *data.FinanceData, monthName string) (*data.Month, error) {
	for i := range financeData.Months {
		if financeData.Months[i].Name == monthName {
			return &financeData.Months[i], nil
		}
	}
	return nil, fmt.Errorf("month %s, does not exist", monthName)
}

func validateDay(day int, month *data.Month) error {
	if day < 1 || day > len(month.Days) {
		return fmt.Errorf("day %d is out of range for month %s", day, month.Name)
	}
	return nil
}
