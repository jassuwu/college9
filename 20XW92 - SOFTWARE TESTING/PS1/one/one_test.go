package one

import (
	"fmt"
	"testing"
)

var cases = []struct{
	Input []int
	Min int
	Max int
}{
	{
		Input: []int{6,9,2,16,19},
		Min: 2, 
		Max: 19, 
	},
	{
		Input: []int{96,11,32,9,39,99,91},
		Min: 9, 
		Max: 99, 
	},
	{
		Input: []int{31,36,42,16,65,76,81},
		Min: 16, 
		Max: 81, 
	},
	{
		Input: []int{28,21,36,31,30,38},
		Min: 21, 
		Max: 38, 
	},
	{
		Input: []int{-10,76,87,-15,12,87},
		Min: -15, 
		Max: 87, 
	},
	{
		Input: []int{6,2,9,5},
		Min: 2, 
		Max: 9, 
	},
	{
		Input: []int{99,21,7},
		Min: 7, 
		Max: 99, 
	},
}

func TestFindMin(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v to %d", test.Input, test.Min), func(t *testing.T) {
			got, want := FindMin(test.Input), test.Min
			assertOutput(t, got, want)
		})
	}
}

func TestFindMax(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%v to %d", test.Input, test.Max), func(t *testing.T) {
			got, want := FindMax(test.Input), test.Max
			assertOutput(t, got, want)
		})
	}
}

func assertOutput(tb testing.TB, got, want int) {
  tb.Helper()
  if got != want {
    tb.Errorf("got %d, want %d", got, want)
  }
}
