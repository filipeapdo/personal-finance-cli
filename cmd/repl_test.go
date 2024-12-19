package cmd

import (
	"strings"
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

func TestHandleAddCommand(t *testing.T) {
	// Create mock finance data
	financeData := data.MockFinanceData()

	// Test case: Valid command
	t.Run("Valid Add Command", func(t *testing.T) {
		input := "add January 1 200.5"
		parts := strings.Fields(input)

		err := handleAddCommand(parts, &financeData)
		if err != nil {
			t.Fatalf("uexpected error: %v", err)
		}

		expected := 310.5
		actual := financeData.Months[0].Days[0].Income
		if actual != expected {
			t.Errorf("income mismatch: expected %.2f, got %.2f", expected, actual)
		}
	})
}
