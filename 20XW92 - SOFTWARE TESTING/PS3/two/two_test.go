package two

import (
	"testing"
)

var cases = []struct {
	Name   string
	Input  int
	Output int
	Error  error
}{
	{
		Name:   "less than 1",
		Input:  0,
		Output: 0,
		Error:  ErrInvalidPrice,
	},
	{
		Name:   "exactly equal to 1",
		Input:  1,
		Output: 0,
		Error:  nil,
	},
	{
		Name:   "within 1 and 49",
		Input:  25,
		Output: 0,
		Error:  nil,
	},
	{
		Name:   "exactly equal to 50",
		Input:  50,
		Output: 0,
		Error:  nil,
	},
	{
		Name:   "within 51 and 200",
		Input:  125,
		Output: 5,
		Error:  nil,
	},
	{
		Name:   "exactly equal to 200",
		Input:  200,
		Output: 5,
		Error:  nil,
	},
	{
		Name:   "within 201 and 500",
		Input:  350,
		Output: 10,
		Error:  nil,
	},
	{
		Name:   "exactly equal to 500",
		Input:  500,
		Output: 10,
		Error:  nil,
	},
	{
		Name:   "greater than 500",
		Input:  501,
		Output: 15,
		Error:  nil,
	},
}

func TestGetDiscountRate(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got, err := GetDiscountRate(test.Input)
			if test.Error == nil {
				want := test.Output
				assertOutput(t, got, want)
				assertNoError(t, err)
			} else {
				assertError(t, err)
			}
		})
	}
}

func assertOutput(tb testing.TB, got, want int) {
	tb.Helper()
	if got != want {
		tb.Errorf("got %d, and want %d", got, want)
	}
}

func assertError(tb testing.TB, err error) {
	tb.Helper()
	if err == nil {
		tb.Errorf("Expected error, but didn't get error.")
	}
}

func assertNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Errorf("Expected no error, but got error.")
	}
}
