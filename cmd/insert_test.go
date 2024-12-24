package cmd

import (
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func TestInsertIncome(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Successful insertition of income
	t.Run("Valid Income Insertion", func(t *testing.T) {
		err := insertAmount(&financeData, "income", "January", 1, 200.5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := 210.5
		actual := financeData.Months[0].Days[0].Income
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		err := insertAmount(&financeData, "income", "FakeMonth", 10, 100.0)
		if err == nil {
			t.Fatal("expected an error for invalid month name, got none")
		}
	})

	// Test case: Invalid day (out of range)
	t.Run("Day Out of Range", func(t *testing.T) {
		err := insertAmount(&financeData, "income", "January", 0, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}

		err = insertAmount(&financeData, "income", "January", 32, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}
	})

	// Test case: Negative day
	t.Run("Negative Day", func(t *testing.T) {
		err := insertAmount(&financeData, "income", "January", -1, 100.0)
		if err == nil {
			t.Fatal("expected an error for negative day, got none")
		}
	})
}

func TestInsertExpense(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Successful insertition of expense
	t.Run("Valid Expense Insertion", func(t *testing.T) {
		err := insertAmount(&financeData, "expense", "January", 1, 15.0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := 23.0
		actual := financeData.Months[0].Days[0].Expense
		if actual != expected {
			t.Errorf("expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		err := insertAmount(&financeData, "expense", "FakeMonth", 10, 100.0)
		if err == nil {
			t.Fatal("expected an error for invalid month name, got none")
		}
	})

	// Test case: Invalid day (out of range)
	t.Run("Day Out of Range", func(t *testing.T) {
		err := insertAmount(&financeData, "expense", "January", 0, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}

		err = insertAmount(&financeData, "expense", "January", 32, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}
	})

	// Test case: Negative day
	t.Run("Negative Day", func(t *testing.T) {
		err := insertAmount(&financeData, "expense", "January", -1, 100.0)
		if err == nil {
			t.Fatal("expected an error for negative day, got none")
		}
	})
}

func TestInsertDailyExpense(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Successful insertition of daily expense
	t.Run("Valid Daily expense Insertion", func(t *testing.T) {
		err := insertAmount(&financeData, "daily", "January", 1, 0.99)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := 1.99
		actual := financeData.Months[0].Days[0].Daily
		if actual != expected {
			t.Errorf("daily expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		err := insertAmount(&financeData, "daily", "FakeMonth", 10, 100.0)
		if err == nil {
			t.Fatal("expected an error for invalid month name, got none")
		}
	})

	// Test case: Invalid day (out of range)
	t.Run("Day Out of Range", func(t *testing.T) {
		err := insertAmount(&financeData, "daily", "January", 0, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}

		err = insertAmount(&financeData, "daily", "January", 32, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}
	})

	// Test case: Negative day
	t.Run("Negative Day", func(t *testing.T) {
		err := insertAmount(&financeData, "daily", "January", -1, 100.0)
		if err == nil {
			t.Fatal("expected an error for negative day, got none")
		}
	})
}
