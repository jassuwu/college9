package four

import (
	"errors"
)

// OrderPizza validates the number of pizzas and returns a corresponding message
func OrderPizza(number int) (string, error) {
	if number < 1 {
		return "", errors.New("Invalid order: Value must be between 1 and 10")
	} else if number > 10 {
		return "", errors.New("Only 10 pizzas can be ordered")
	} else {
		return "Order placed successfully", nil
	}
}
