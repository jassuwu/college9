package three

import (
	"errors"
)

var (
	ErrInputOutOfRange = errors.New("input values out of range")
	ErrInvalidTriangle = errors.New("invalid triangle")
)

// ClassifyTriangle classifies a triangle based on side lengths a, b, and c.
func ClassifyTriangle(a, b, c int) (string, error) {
	// Check for invalid input
	if a <= 0 || b <= 0 || c <= 0 || a > 100 || b > 100 || c > 100 {
		return "", ErrInputOutOfRange
	}

	// Sort sides to simplify classification
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	// Check for valid triangle
	if a+b <= c {
		return "", ErrInvalidTriangle
	}

	// Classification
	a2 := a * a
	b2 := b * b
	c2 := c * c

	if c2 == a2+b2 {
		return "Right angled triangle", nil
	} else if c2 > a2+b2 {
		return "Obtuse angled triangle", nil
	} else if c2 < a2+b2 {
		return "Acute angled triangle", nil
	}

	return "", nil
}
