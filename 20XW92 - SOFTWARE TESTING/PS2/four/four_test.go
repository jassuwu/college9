package four

import (
	"math"
	"testing"
)

func TestCalculateSampleStandardDeviation(t *testing.T) {
	tests := []struct {
		Name        string
		Sample      []float64
		Expected    float64
		ExpectedErr error
	}{
		// Valid cases
		{
			Name:        "Normal Case",
			Sample:      []float64{1, 2, 3, 4, 5},
			Expected:    1.5811388300841898, // Calculated manually or using a reliable tool
			ExpectedErr: nil,
		},
		{
			Name:        "Small Sample Size",
			Sample:      []float64{10, 12},
			Expected:    1.4142135623730951, // Calculated manually or using a reliable tool
			ExpectedErr: nil,
		},

		// Edge cases
		{
			Name:        "Single Element Sample",
			Sample:      []float64{5},
			Expected:    0,
			ExpectedErr: ErrInvalidSampleSize,
		},
		{
			Name:        "Two Elements",
			Sample:      []float64{3, 3},
			Expected:    0,
			ExpectedErr: ErrInvalidSampleSize,
		},

		// Invalid cases
		{
			Name:        "Empty Sample",
			Sample:      []float64{},
			Expected:    0,
			ExpectedErr: ErrInvalidSampleSize,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := CalculateSampleStandardDeviation(tt.Sample)
			if err != nil && err.Error() != tt.ExpectedErr.Error() {
				t.Errorf(
					"CalculateSampleStandardDeviation() error = %v, wantErr %v",
					err,
					tt.ExpectedErr,
				)
				return
			}
			if math.Abs(got-tt.Expected) > 1e-9 {
				t.Errorf("CalculateSampleStandardDeviation() = %v, want %v", got, tt.Expected)
			}
		})
	}
}
