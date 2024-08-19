package four

import (
	"errors"
	"math"
)

var ErrInvalidSampleSize = errors.New("sample size must be at least 2")

// CalculateSampleStandardDeviation calculates the standard deviation of a sample.
func CalculateSampleStandardDeviation(sample []float64) (float64, error) {
	n := len(sample)
	if n < 2 {
		return 0, ErrInvalidSampleSize
	}

	var sum float64
	for _, value := range sample {
		sum += value
	}
	mean := sum / float64(n)

	var varianceSum float64
	for _, value := range sample {
		diff := value - mean
		varianceSum += diff * diff
	}

	// Sample standard deviation formula: sqrt(sum of squared differences / (n - 1))
	return math.Sqrt(varianceSum / float64(n-1)), nil
}
