package data

func MockFinanceData() FinanceData {
	monthNames := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	var months []Month
	for _, monthName := range monthNames {
		var days []Day
		for day := 1; day <= 30; day++ {
			income := 10.0
			expense := 8.0
			daily := 1.0

			days = append(days, Day{
				Day:     day,
				Income:  income,
				Expense: expense,
				Daily:   daily,
			})
		}

		var totalIncome, totalExpense, totalDaily float64
		for _, day := range days {
			totalIncome += day.Income
			totalExpense += day.Expense
			totalDaily += day.Daily
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
			},
		})
	}

	return FinanceData{Months: months, FilePath: "test_finance_data.json"}
}
