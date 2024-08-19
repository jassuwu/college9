package three

import (
	"errors"
)

var (
	ErrInvalidMarks = errors.New("Marks should be in the range 0-100.")
)

func CalculateGrade(mark1, mark2, mark3 int) (string, error) {
	if mark1 < 0 || mark1 > 100 {
		return "", ErrInvalidMarks
	}
	if mark2 < 0 || mark2 > 100 {
		return "", ErrInvalidMarks
	}
	if mark3 < 0 || mark3 > 100 {
		return "", ErrInvalidMarks
	}

	totalMarks := (mark1 + mark2 + mark3) / 3

	if totalMarks >= 90 && totalMarks <= 100 {
		return "First class Distinction", nil
	} else if totalMarks >= 75 && totalMarks <= 89 {
		return "First Class", nil
	} else if totalMarks >= 60 && totalMarks <= 74 {
		return "Second Class", nil
	} else if totalMarks >= 50 && totalMarks <= 59 {
		return "Third Class", nil
	} else if totalMarks >= 0 && totalMarks < 50 {
		return "Fail", nil
	} else {
		return "", ErrInvalidMarks
	}
}
