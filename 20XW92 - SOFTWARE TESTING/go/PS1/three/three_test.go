package three

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected *big.Int
		err      error
	}{
		{0, big.NewInt(1), nil},
		{1, big.NewInt(1), nil},
		{5, big.NewInt(120), nil},
		{10, big.NewInt(3628800), nil},
		{20, big.NewInt(2432902008176640000), nil},
		{-1, nil, ErrInvalidInput},
	}

	for _, test := range tests {
		result, err := Factorial(test.input)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("For input %d, expected error %v but got %v", test.input, test.err, err)
		}
		if result != nil && result.Cmp(test.expected) != 0 {
			t.Errorf("For input %d, expected %s but got %s", test.input, test.expected, result)
		}
	}
}
