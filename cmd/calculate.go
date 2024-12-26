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

func calculateSummary(fd *data.FinanceData) {
	for monthIndex := range fd.Months {
		month := &fd.Months[monthIndex]
		month.Summary.TotalIncome = 0
		month.Summary.TotalExpense = 0
		month.Summary.TotalDaily = 0
		month.Summary.FinalBalance = 0
		for _, day := range month.Days {
			month.Summary.TotalIncome += day.Income
			month.Summary.TotalExpense += day.Expense
			month.Summary.TotalDaily += day.Daily
		}
		if len(month.Days) > 0 {
			month.Summary.FinalBalance = month.Days[len(month.Days)-1].Balance
		}
	}
}
