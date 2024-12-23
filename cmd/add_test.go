package cmd

import (
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func TestAddIncome(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Successful addition of income
	t.Run("Valid Income Addition", func(t *testing.T) {
		err := addAmount(&financeData, "income", "January", 1, 200.5)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := 310.5
		actual := financeData.Months[0].Days[0].Income
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		err := addAmount(&financeData, "income", "FakeMonth", 10, 100.0)
		if err == nil {
			t.Fatal("expected an error for invalid month name, got none")
		}
	})

	// Test case: Invalid day (out of range)
	t.Run("Day Out of Range", func(t *testing.T) {
		err := addAmount(&financeData, "income", "January", 0, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}

		err = addAmount(&financeData, "income", "January", 32, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}
	})

	// Test case: Negative day
	t.Run("Negative Day", func(t *testing.T) {
		err := addAmount(&financeData, "income", "January", -1, 100.0)
		if err == nil {
			t.Fatal("expected an error for negative day, got none")
		}
	})
}

func TestAddExpense(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Successful addition of expense
	t.Run("Valid Expense Addition", func(t *testing.T) {
		err := addAmount(&financeData, "expense", "January", 1, 15.0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := 70.0
		actual := financeData.Months[0].Days[0].Expense
		if actual != expected {
			t.Errorf("expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		err := addAmount(&financeData, "expense", "FakeMonth", 10, 100.0)
		if err == nil {
			t.Fatal("expected an error for invalid month name, got none")
		}
	})

	// Test case: Invalid day (out of range)
	t.Run("Day Out of Range", func(t *testing.T) {
		err := addAmount(&financeData, "expense", "January", 0, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}

		err = addAmount(&financeData, "expense", "January", 32, 100.0)
		if err == nil {
			t.Fatal("expected an error for day out of range, got none")
		}
	})

	// Test case: Negative day
	t.Run("Negative Day", func(t *testing.T) {
		err := addAmount(&financeData, "expense", "January", -1, 100.0)
		if err == nil {
			t.Fatal("expected an error for negative day, got none")
		}
	})
}
