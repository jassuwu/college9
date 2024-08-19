package four

import (
	"errors"
)

var (
	ErrInvalidOrder = errors.New("Invalid order: Value must be between 1 and 10")
	ErrMorePizzas = errors.New("Only 10 pizzas can be ordered")
)
// OrderPizza validates the number of pizzas and returns a corresponding message
func OrderPizza(number int) (string, error) {
	if number < 1 {
		return "", ErrInvalidOrder
	} else if number > 10 {
		return "", ErrMorePizzas 
	} else {
		return "Order placed successfully", nil
	}
}
