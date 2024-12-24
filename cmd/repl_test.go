package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func TestHandleInsertCommand(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Schedule cleanup to remove test_finance_data.json
	t.Cleanup(func() {
		err := os.Remove(financeData.FilePath)
		if err != nil {
			t.Errorf("failed to delete test_finance_data.json: %v", err)
		}
	})

	// Test case: Valid insert income command
	t.Run("Valid Insert Income Command", func(t *testing.T) {
		input := "insert income January 1 200.5"
		parts := strings.Fields(input)

		err := handleInsertCommand(&financeData, parts)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 210.5
		actual := financeData.Months[0].Days[0].Income
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Valid insert expense command
	t.Run("Valid Insert Expense Command", func(t *testing.T) {
		input := "insert expense January 1 15.0"
		parts := strings.Fields(input)

		err := handleInsertCommand(&financeData, parts)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 23.0
		actual := financeData.Months[0].Days[0].Expense
		if actual != expected {
			t.Errorf("expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})

	// Test case: Valid insert daily expense command
	t.Run("Valid Insert Daily expense Command", func(t *testing.T) {
		input := "insert daily January 1 0.99"
		parts := strings.Fields(input)

		err := handleInsertCommand(&financeData, parts)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 1.99
		actual := financeData.Months[0].Days[0].Daily
		if actual != expected {
			t.Errorf("daily expense mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})
}
