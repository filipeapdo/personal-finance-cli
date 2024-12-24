package cmd

import (
	"testing"

	"github.com/filipeapdo/personal-finance-cli/data"
)

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
