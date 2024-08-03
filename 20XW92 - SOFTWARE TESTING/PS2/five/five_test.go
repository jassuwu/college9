package five

import (
	"testing"
)

func TestGetDayOfWeek(t *testing.T) {
	tests := []struct {
		Name        string
		Day         int
		Month       int
		Year        int
		Expected    string
		ExpectedErr error
	}{
		// Valid dates
		{"Valid Date - Wednesday", 14, 8, 2002, "Wednesday", nil},
		{"Valid Date - Monday", 1, 1, 2001, "Monday", nil},
		{"Valid Date - Friday", 31, 12, 2010, "Friday", nil},

		// Invalid dates
		{"Invalid Day", 32, 1, 2000, "", ErrInvalidDateRange},
		{"Invalid Month", 15, 13, 2000, "", ErrInvalidDateRange},
		{"Invalid Year", 15, 1, 2021, "", ErrInvalidDateRange},
		{"Invalid Date", 30, 2, 2000, "", ErrInvalidDate},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := GetDayOfWeek(tt.Day, tt.Month, tt.Year)
			if err != nil && err.Error() != tt.ExpectedErr.Error() {
				t.Errorf("GetDayOfWeek() error = %v, wantErr %v", err, tt.ExpectedErr)
				return
			}
			if got != tt.Expected {
				t.Errorf("GetDayOfWeek() = %v, want %v", got, tt.Expected)
			}
		})
	}
}
