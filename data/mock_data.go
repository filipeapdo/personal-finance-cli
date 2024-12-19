package data

func MockFinanceData() FinanceData {
	monthNames := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	var months []Month
	for i, monthName := range monthNames {
		var days []Day
		for day := 1; day <= 30; day++ {
			income := float64((i+1)*100 + day*10)
			expense := float64((i+1)*50 + day*5)
			daily := float64(20)
			balance := income - expense - daily

			days = append(days, Day{
				Day:     day,
				Income:  income,
				Expense: expense,
				Daily:   daily,
				Balance: balance,
			})
		}

		var totalIncome, totalExpense, totalDaily, finalBalance float64
		for _, day := range days {
			totalIncome += day.Income
			totalExpense += day.Expense
			totalDaily += day.Daily
			finalBalance = day.Balance
		}

		months = append(months, Month{
			Name: monthName,
			Days: days,
			Summary: struct {
				TotalIncome  float64 `json:"total_income"`
				TotalExpense float64 `json:"total_expense"`
				TotalDaily   float64 `json:"total_daily"`
				FinalBalance float64 `json:"final_balance"`
			}{
				TotalIncome:  totalIncome,
				TotalExpense: totalExpense,
				TotalDaily:   totalDaily,
				FinalBalance: finalBalance,
			},
		})
	}

	return FinanceData{Months: months}
}
