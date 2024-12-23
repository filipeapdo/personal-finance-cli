package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func TestHandleAddCommand(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Schedule cleanup to remove test_finance_data.json
	t.Cleanup(func() {
		err := os.Remove(financeData.FilePath)
		if err != nil {
			t.Errorf("failed to delete test_finance_data.json: %v", err)
		}
	})

	// Test case: Valid add income command
	t.Run("Valid Add Income Command", func(t *testing.T) {
		input := "add income January 1 200.5"
		parts := strings.Fields(input)

		err := handleAddCommand(&financeData, parts)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 310.5
		actual := financeData.Months[0].Days[0].Income
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Valid add expense command
	t.Run("Valid Add Expense Command", func(t *testing.T) {
		input := "add expense January 1 15.0"
		parts := strings.Fields(input)

		err := handleAddCommand(&financeData, parts)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 70.0
		actual := financeData.Months[0].Days[0].Expense
		if actual != expected {
			t.Errorf("expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})
}
