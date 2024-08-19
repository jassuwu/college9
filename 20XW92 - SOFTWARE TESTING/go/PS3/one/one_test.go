package one

import (
	"testing"
)

var cases = []struct{
	Name string
	Input string
	Output bool
}{
	{
		Name: "length less than 6",
		Input: "pass",
		Output: false,
	},
	{
		Name: "exactly equal to 6",
		Input: "passwd",
		Output: true,
	},
	{
		Name: "between 7 and 11",
		Input: "password",
		Output: true,
	},
	{
		Name: "exactly equal to 12",
		Input: "passwordishy",
		Output: true,
	},
	{
		Name: "length greater than 12",
		Input: "passwordyword",
		Output: false,
	},
}

func TestValidPasswordLength(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got, want := ValidPasswordLength(test.Input), test.Output
			assertOutput(t, got, want)
		})
	}
}

func assertOutput(tb testing.TB, got, want bool) {
	tb.Helper()
	if got != want {
		tb.Errorf("got %v, want %v", got, want)
	}
}
