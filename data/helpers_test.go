package data

import (
	"reflect"
	"testing"
)

func TestDataHelpersSorts(t *testing.T) {
	// Test case: Sort FinanceData Correctly
	t.Run("Sort FinanceData Correctly", func(t *testing.T) {
		mockData := FinanceData{
			Months: []Month{
				{Name: "March", Days: []Day{{Day: 5}, {Day: 3}, {Day: 1}}},
				{Name: "January", Days: []Day{{Day: 2}, {Day: 1}, {Day: 3}}},
				{Name: "February", Days: []Day{{Day: 28}, {Day: 14}, {Day: 1}}},
			},
		}

		expectedData := FinanceData{
			Months: []Month{
				{Name: "January", Days: []Day{{Day: 1}, {Day: 2}, {Day: 3}}},
				{Name: "February", Days: []Day{{Day: 1}, {Day: 14}, {Day: 28}}},
				{Name: "March", Days: []Day{{Day: 1}, {Day: 3}, {Day: 5}}},
			},
		}

		SortFinanceData(&mockData)
		if !reflect.DeepEqual(mockData, expectedData) {
			t.Errorf("unexpected error sorting data:\nexpected %+v,\ngot %+v", expectedData, mockData)
		}
	})

	// Test case: Edge Cases: Empty FinanceData
	t.Run("Edge Case: Empty FinanceData", func(t *testing.T) {
		mockData := FinanceData{}
		expectedData := FinanceData{}

		SortFinanceData(&mockData)
		if !reflect.DeepEqual(mockData, expectedData) {
			t.Errorf("unexpected error for edge case: 'empty data':\nexpected %+v,\ngot %+v", expectedData, mockData)
		}
	})

	// Test case: Edge Cases: Single Month with No Days
	t.Run("Edge Case: Single Month with No Days", func(t *testing.T) {
		mockData := FinanceData{
			Months: []Month{
				{Name: "April", Days: []Day{}},
			},
		}
		expectedData := mockData

		SortFinanceData(&mockData)
		if !reflect.DeepEqual(mockData, expectedData) {
			t.Errorf("unexpected error for edge case: 'single month with no days':\nexpected %+v,\ngot %+v", expectedData, mockData)
		}
	})

	// Test case: Edge Cases: Already Sorted FinanceData
	t.Run("Edge Case: Already Sorted Data", func(t *testing.T) {
		mockData := FinanceData{
			Months: []Month{
				{Name: "January", Days: []Day{{Day: 1}, {Day: 2}, {Day: 3}}},
				{Name: "February", Days: []Day{{Day: 1}, {Day: 14}, {Day: 28}}},
				{Name: "March", Days: []Day{{Day: 1}, {Day: 3}, {Day: 5}}},
			},
		}
		expectedData := mockData

		SortFinanceData(&mockData)
		if !reflect.DeepEqual(mockData, expectedData) {
			t.Errorf("unexpected error for edge case: 'already sorted data':\nexpected %+v,\ngot %+v", expectedData, mockData)
		}
	})
}

func TestDataHelpersValidations(t *testing.T) {
	// Test case: Valid month name
	t.Run("Valid Month Name", func(t *testing.T) {
		validMonths := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
		for _, month := range validMonths {
			err := ValidateMonth(month)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
	})

	// Test case: Invalid month name
	t.Run("Invalid Month Name", func(t *testing.T) {
		invalidMonths := []string{"FakeMonth", "Jan", "", "Month123"}
		for _, month := range invalidMonths {
			err := ValidateMonth(month)
			if err == nil {
				t.Fatalf("expected an error for invalid month name, got none")
			}
		}
	})

	// Test case: Valid day - 31 days: Jan, Mar, May, Jul, Aug, Oct, Dec
	t.Run("Valid Days - Months with 31 days", func(t *testing.T) {
		monthWith31Days := []string{"January", "March", "May", "July", "August", "October", "December"}
		for _, month := range monthWith31Days {
			for day := 1; day <= 31; day++ {
				err := ValidateDay(month, day, 2024)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			// Teste invalid day
			err := ValidateDay(month, 32, 2024)
			if err == nil {
				t.Fatalf("expected an error for %s day 32, got none", month)
			}
		}
	})

	// Test case: Valid day - 30 days: Apr, Jun, Sep, Nov
	t.Run("Valid Days - Months with 30 days", func(t *testing.T) {
		monthWith30Days := []string{"April", "June", "September", "November"}
		for _, month := range monthWith30Days {
			for day := 1; day <= 30; day++ {
				err := ValidateDay(month, day, 2024)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			// Teste invalid day
			err := ValidateDay(month, 31, 2024)
			if err == nil {
				t.Fatalf("expected an error for %s day 31, got none", month)
			}
		}
	})

	// Test case: Valid and Invalid days for February
	t.Run("Valid Days - February", func(t *testing.T) {
		// Non-lear year
		for day := 1; day <= 28; day++ {
			err := ValidateDay("February", day, 2023)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		// Test invalid day
		err := ValidateDay("February", 29, 2023)
		if err == nil {
			t.Fatalf("expected an error for February day 29 in non-leap year, got none")
		}

		// Leap year
		for day := 1; day <= 29; day++ {
			err := ValidateDay("February", day, 2024)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		}
		// Test invalid day
		err = ValidateDay("February", 30, 2024)
		if err == nil {
			t.Fatalf("expected an error for February day 30 in leap year, got none")
		}
	})

	// Test case: Edge Cases
	t.Run("Edge Cases", func(t *testing.T) {
		edgeCases := []struct {
			month string
			day   int
			year  int
		}{
			{"January", 0, 2024},    // Invalid Day
			{"January", -1, 2024},   // Negative Day
			{"", 15, 2024},          // Empty Month
			{"FakeMonth", 15, 2024}, // Invalid Month
		}

		for _, tc := range edgeCases {
			err := ValidateDay(tc.month, tc.day, tc.year)
			if err == nil {
				t.Fatalf("expected an error for %s day %d year %d, got none", tc.month, tc.day, tc.year)
			}
		}
	})
}
