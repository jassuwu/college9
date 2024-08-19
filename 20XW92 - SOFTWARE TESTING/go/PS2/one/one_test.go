package one

import (
	"testing"
)

var cases = []struct {
	Name       string
	Input      string
	VowelCount int
	TotalCount int
}{
	{Name: "Empty character", Input: "", VowelCount: 0, TotalCount: 0},
	{Name: "Single character w/o vowel", Input: "b", VowelCount: 0, TotalCount: 1},
	{Name: "Single character w/ vowel", Input: "a", VowelCount: 1, TotalCount: 1},
	{Name: "Multiple vowels", Input: "aeiou", VowelCount: 5, TotalCount: 5},
	{Name: "Multiple consonants", Input: "bcdfg", VowelCount: 0, TotalCount: 5},
	{Name: "Mixed characters", Input: "abcde12345", VowelCount: 2, TotalCount: 5},
	{Name: "Non-alphabetic characters only", Input: "1234567890", VowelCount: 0, TotalCount: 0},
	{Name: "Mixed with non-alphabetic and spaces", Input: "a b c 1 2 e", VowelCount: 2, TotalCount: 4},
	{Name: "Long string with no vowels", Input: "bcdf" + string([]rune{'b', 'c', 'd', 'f'}), VowelCount: 0, TotalCount: 8},
	{Name: "Long string with only vowels", Input: "aeiou" + string([]rune{'a', 'e', 'i', 'o', 'u'}), VowelCount: 10, TotalCount: 10},
	{Name: "String with uppercase letters", Input: "AEIOUaeiou", VowelCount: 10, TotalCount: 10},
	{Name: "String with mixed case and non-alphabetic", Input: "A1e2I3o4U5", VowelCount: 5, TotalCount: 5},
}

func TestCountChar(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			gotVCount, gotTCount := CountChar(test.Input)
			assertOutput(t, [2]int{gotVCount, gotTCount}, [2]int{test.VowelCount, test.TotalCount})
		})
	}
}

func assertOutput(tb testing.TB, got [2]int, want [2]int) {
	if got[0] != want[0] || got[1] != want[1] {
		tb.Errorf("Failed. Got %v, want %v", got, want)
	}
}
