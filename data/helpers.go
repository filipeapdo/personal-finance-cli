package data

import (
	"fmt"
	"sort"
)

var months = map[string]struct {
	Order   int
	MaxDays int
}{
	"January":   {1, 31},
	"February":  {2, 28},
	"March":     {3, 31},
	"April":     {4, 30},
	"May":       {5, 31},
	"June":      {6, 30},
	"July":      {7, 31},
	"August":    {8, 31},
	"September": {9, 30},
	"October":   {10, 31},
	"November":  {11, 30},
	"December":  {12, 31},
}

func SortFinanceData(fd *FinanceData) {
	sortMonths(fd)
	for i := range fd.Months {
		sortDays(&fd.Months[i])
	}
}

func sortMonths(fd *FinanceData) {
	sort.Slice(fd.Months, func(i, j int) bool {
		return months[fd.Months[i].Name].Order < months[fd.Months[j].Name].Order
	})
}

func sortDays(month *Month) {
	sort.Slice(month.Days, func(i, j int) bool {
		return month.Days[i].Day < month.Days[j].Day
	})
}

func ValidateMonth(monthName string) error {
	_, exists := months[monthName]
	if !exists {
		return fmt.Errorf("invalid month name, valid month's names are: %v", months)
	}

	return nil
}

func ValidateDay(monthName string, day, year int) error {
	var maxDay int
	err := ValidateMonth(monthName)
	if err != nil {
		return err
	}

	maxDay = months[monthName].MaxDays
	if monthName == "February" {
		if isLeapYear(year) {
			maxDay = 29
		}
	}

	if day < 1 || day > maxDay {
		return fmt.Errorf("invalid day, valid days for %s are: between 1 and %d", monthName, maxDay)
	}

	return nil
}

// A year is a leap year if it is divisible by 4, but not divisible by 100, unless it is also divisible by 400
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
