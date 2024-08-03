package five

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidDateRange = errors.New("date out of valid range")
	ErrInvalidDate      = errors.New("invalid date")
)

// GetDayOfWeek returns the day of the week for a given date.
func GetDayOfWeek(day, month, year int) (string, error) {
	// Validate inputs
	if day < 1 || day > 31 || month < 1 || month > 12 || year < 2000 || year > 2019 {
		return "", ErrInvalidDateRange
	}

	// Check if the date is valid
	t, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-%02d", year, month, day))
	if err != nil {
		return "", ErrInvalidDate
	}

	// Return the day of the week
	return t.Weekday().String(), nil
}
