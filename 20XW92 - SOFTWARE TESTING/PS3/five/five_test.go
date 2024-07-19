package one

import (
	"testing"
)

var taxCases = []struct {
	Name   string
	Input  int
	Output float64
	Err    error
}{
	{
		Name:   "pay less than 12000",
		Input:  11000,
		Output: 0,
		Err:    ErrInvalidPay,
	},
	{
		Name:   "pay exactly 12000",
		Input:  12000,
		Output: 0,
		Err:    nil,
	},
	{
		Name:   "pay between 12000 and 15000",
		Input:  13000,
		Output: 0,
		Err:    nil,
	},
	{
		Name:   "pay exactly 15000",
		Input:  15000,
		Output: 0,
		Err:    nil,
	},
	{
		Name:   "pay between 15001 and 25000",
		Input:  20000,
		Output: 3600,
		Err:    nil,
	},
	{
		Name:   "pay exactly 25000",
		Input:  25000,
		Output: 4500,
		Err:    nil,
	},
	{
		Name:   "pay between 25001 and 35000",
		Input:  30000,
		Output: 6000,
		Err:    nil,
	},
	{
		Name:   "pay exactly 35000",
		Input:  35000,
		Output: 7000,
		Err:    nil,
	},
	{
		Name:   "pay greater than 35000",
		Input:  36000,
		Output: 0,
		Err:    ErrInvalidPay,
	},
}

func TestCalculateTax(t *testing.T) {
	for _, test := range taxCases {
		t.Run(test.Name, func(t *testing.T) {
			got, err := CalculateTax(test.Input)
			if got != test.Output || (err != nil && err.Error() != test.Err.Error()) {
				t.Errorf(
					"CalculateTax(%d) = %v, %v; want %v, %v",
					test.Input,
					got,
					err,
					test.Output,
					test.Err,
				)
			}
		})
	}
}
