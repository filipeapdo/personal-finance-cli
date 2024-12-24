package cmd

import (
	"github.com/filipeapdo/personal-finance-cli/data"
)

func calculateBalance(fd *data.FinanceData) {
	previousBalance := 0.0
	for monthIndex, month := range fd.Months {
		for dayIndex := range month.Days {
			day := &fd.Months[monthIndex].Days[dayIndex]
			if monthIndex == 0 && dayIndex == 0 {
				day.Balance = day.Income - day.Expense - day.Daily
			} else {
				day.Balance = previousBalance + day.Income - day.Expense - day.Daily
			}
			previousBalance = day.Balance
		}
	}
}
