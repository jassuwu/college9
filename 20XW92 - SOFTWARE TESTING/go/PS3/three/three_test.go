package three

import (
	"testing"
)

var gradeCases = []struct {
	Name   string // Test case name
	Mark1  int    // Marks for subject 1
	Mark2  int    // Marks for subject 2
	Mark3  int    // Marks for subject 3
	Grade  string // Expected grade
	Err    error  // Expected error
}{
	{
		Name:   "valid marks for first class distinction",
		Mark1:  95,
		Mark2:  98,
		Mark3:  92,
		Grade:  "First class Distinction",
		Err:    nil,
	},
	{
		Name:   "valid marks for first class",
		Mark1:  85,
		Mark2:  80,
		Mark3:  89,
		Grade:  "First Class",
		Err:    nil,
	},
	{
		Name:   "valid marks for second class",
		Mark1:  70,
		Mark2:  65,
		Mark3:  72,
		Grade:  "Second Class",
		Err:    nil,
	},
	{
		Name:   "valid marks for third class",
		Mark1:  55,
		Mark2:  56,
		Mark3:  58,
		Grade:  "Third Class",
		Err:    nil,
	},
	{
		Name:   "valid marks for fail",
		Mark1:  30,
		Mark2:  25,
		Mark3:  40,
		Grade:  "Fail",
		Err:    nil,
	},
	{
		Name:   "invalid marks (negative marks)",
		Mark1:  -10,
		Mark2:  60,
		Mark3:  75,
		Grade:  "",
		Err:    ErrInvalidMarks,
	},
	{
		Name:   "invalid marks (marks exceeding 100)",
		Mark1:  110,
		Mark2:  95,
		Mark3:  85,
		Grade:  "",
		Err:    ErrInvalidMarks,
	},
}

func TestCalculateGrade(t *testing.T) {
	for _, test := range gradeCases {
		t.Run(test.Name, func(t *testing.T) {
			gotGrade, gotErr := CalculateGrade(test.Mark1, test.Mark2, test.Mark3)

			// Check if error matches expected error
			if (gotErr != nil && test.Err == nil) || (gotErr == nil && test.Err != nil) || (gotErr != nil && test.Err != nil && gotErr.Error() != test.Err.Error()) {
				t.Errorf("CalculateGrade(%d, %d, %d) error = %v, want %v", test.Mark1, test.Mark2, test.Mark3, gotErr, test.Err)
			}

			// Check if grade matches expected grade
			if gotGrade != test.Grade {
				t.Errorf("CalculateGrade(%d, %d, %d) = %s, want %s", test.Mark1, test.Mark2, test.Mark3, gotGrade, test.Grade)
			}
		})
	}
}

