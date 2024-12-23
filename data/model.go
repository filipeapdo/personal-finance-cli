package data

type Day struct {
	Day     int     `json:"day"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Daily   float64 `json:"daily"`
	Balance float64 `json:"balance"`
}

type Month struct {
	Name    string `json:"name"`
	Days    []Day  `json:"days"`
	Summary struct {
		TotalIncome  float64 `json:"total_income"`
		TotalExpense float64 `json:"total_expense"`
		TotalDaily   float64 `json:"total_daily"`
		FinalBalance float64 `json:"final_balance"`
	} `json:"summary"`
}

type FinanceData struct {
	Months   []Month `json:"months"`
	FilePath string  `json:"filepath"`
}
