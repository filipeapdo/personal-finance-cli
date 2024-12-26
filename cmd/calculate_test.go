package cmd

import (
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

// to-do: review tests with meaningful mockdata (?)

func TestCalculateBalance(t *testing.T) {
	financeData := data.MockFinanceData()

	// Test case: Successful balance calculated
	t.Run("Valid balance calculated", func(t *testing.T) {
		calculateBalance(&financeData)

		// January 1
		expected := 1.0
		actual := financeData.Months[0].Days[0].Balance
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}

		// January 10
		expected = 10.0
		actual = financeData.Months[0].Days[9].Balance
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}

		// December 30
		expected = 360.0
		actual = financeData.Months[11].Days[29].Balance
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})
}

func TestCalculateSummary(t *testing.T) {
	financeData := data.MockFinanceData()

	// Test case: Successful summary calculated
	t.Run("Valid summary calculated", func(t *testing.T) {
		calculateBalance(&financeData)
		calculateSummary(&financeData)

		// January
		expectedTotalIncome := 300.0
		expectedTotalExpense := 240.0
		expectedTotalDaily := 30.0
		expectedFinalBalance := 30.0

		actualTotalIncome := financeData.Months[0].Summary.TotalIncome
		actualTotalExpense := financeData.Months[0].Summary.TotalExpense
		actualTotalDaily := financeData.Months[0].Summary.TotalDaily
		actualFinalBalance := financeData.Months[0].Summary.FinalBalance

		if expectedTotalIncome != actualTotalIncome {
			t.Errorf("Total Income mismatch: expected %.2f, got %.2f", expectedTotalIncome, actualTotalIncome)
		}
		if expectedTotalExpense != actualTotalExpense {
			t.Errorf("Total Expense mismatch: expected %.2f, got %.2f", expectedTotalExpense, actualTotalExpense)
		}
		if expectedTotalDaily != actualTotalDaily {
			t.Errorf("Total Daily mismatch: expected %.2f, got %.2f", expectedTotalDaily, actualTotalDaily)
		}
		if expectedFinalBalance != actualFinalBalance {
			t.Errorf("Total Balance mismatch: expected %.2f, got %.2f", expectedFinalBalance, actualFinalBalance)
		}
	})
}
