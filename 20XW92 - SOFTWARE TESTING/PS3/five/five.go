package one

import "errors"

var ErrInvalidPay = errors.New("Invalid pay: Must be between 12000 and 35000")

// CalculateTax computes the tax based on the given pay
func CalculateTax(pay int) (float64, error) {
	if pay < 12000 || pay > 35000 {
		return 0, ErrInvalidPay
	}

	if pay <= 15000 {
		return 0, nil
	} else if pay <= 25000 {
		return float64(pay) * 0.18, nil
	} else {
		return float64(pay) * 0.20, nil
	}
}
