package three

import (
	"testing"
)

var testCases = []struct {
	Name        string
	A           int
	B           int
	C           int
	Expected    string
	ExpectedErr error
}{
	// Valid triangles
	{
		Name:        "Right angled triangle",
		A:           3,
		B:           4,
		C:           5,
		Expected:    "Right angled triangle",
		ExpectedErr: nil,
	},
	{
		Name:        "Obtuse angled triangle",
		A:           5,
		B:           7,
		C:           10,
		Expected:    "Obtuse angled triangle",
		ExpectedErr: nil,
	},
	{
		Name:        "Acute angled triangle",
		A:           6,
		B:           7,
		C:           8,
		Expected:    "Acute angled triangle",
		ExpectedErr: nil,
	},

	// Edge cases
	{
		Name:        "Sum of two sides equals third (right)",
		A:           3,
		B:           4,
		C:           7,
		Expected:    "",
		ExpectedErr: ErrInvalidTriangle,
	},
	{
		Name:        "Sum of two sides equals third (obtuse)",
		A:           5,
		B:           5,
		C:           10,
		Expected:    "",
		ExpectedErr: ErrInvalidTriangle,
	},

	// Invalid input values
	{
		Name:        "Side with zero value",
		A:           0,
		B:           5,
		C:           7,
		Expected:    "",
		ExpectedErr: ErrInputOutOfRange,
	},
	{
		Name:        "Side with negative value",
		A:           -3,
		B:           4,
		C:           5,
		Expected:    "",
		ExpectedErr: ErrInputOutOfRange,
	},
	{
		Name:        "Value out of range",
		A:           101,
		B:           4,
		C:           5,
		Expected:    "",
		ExpectedErr: ErrInputOutOfRange,
	},

	// Non-integer values (invalid scenario, not directly testable in Go as Go is a strongly typed language)
	// The test should be adapted if inputs were strings or other types, but for integers, this is not applicable.

	// Wrong number of values (this is more about input parsing, not covered by our Go function directly)
	// These scenarios should be handled by input validation outside of this function or adjusted in a broader context.

}

func TestClassifyTriangle(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got, err := ClassifyTriangle(tc.A, tc.B, tc.C)
			if got != tc.Expected || (err != nil && err.Error() != tc.ExpectedErr.Error()) {
				t.Errorf(
					"ClassifyTriangle(%d, %d, %d) = (%q, %v); want (%q, %v)",
					tc.A,
					tc.B,
					tc.C,
					got,
					err,
					tc.Expected,
					tc.ExpectedErr,
				)
			}
		})
	}
}
