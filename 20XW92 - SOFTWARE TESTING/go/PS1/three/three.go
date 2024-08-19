package three

import (
	"errors"
	"math/big"
)

var (
	ErrInvalidInput = errors.New("factorial of a negative number is not defined")
)

// Factorial calculates the factorial of a non-negative integer n.
// It returns an error if n is negative.
func Factorial(n int) (*big.Int, error) {
	if n < 0 {
		return nil, ErrInvalidInput
	}

	// Use big.Int for large values
	result := big.NewInt(1)

	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result, nil
}
