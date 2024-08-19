package four

import (
	"testing"
)

var pizzaCases = []struct {
	Name   string
	Input  int
	Output string
	Err    error
}{
	{
		Name:   "order less than 1",
		Input:  0,
		Output: "",
		Err:   ErrInvalidOrder,
	},
	{
		Name:   "order exactly 1",
		Input:  1,
		Output: "Order placed successfully",
		Err:    nil,
	},
	{
		Name:   "order between 1 and 10",
		Input:  5,
		Output: "Order placed successfully",
		Err:    nil,
	},
	{
		Name:   "order exactly 10",
		Input:  10,
		Output: "Order placed successfully",
		Err:    nil,
	},
	{
		Name:   "order more than 10",
		Input:  11,
		Output: "",
		Err:    ErrMorePizzas,
	},
	{
		Name:   "order with three-digit number",
		Input:  100,
		Output: "",
		Err:   ErrMorePizzas,
	},
}

func TestOrderPizza(t *testing.T) {
	for _, test := range pizzaCases {
		t.Run(test.Name, func(t *testing.T) {
			got, err := OrderPizza(test.Input)
			if got != test.Output || (err != nil && err.Error() != test.Err.Error()) {
				t.Errorf(
					"OrderPizza(%d) = %v, %v; want %v, %v",
					test.Input,
					got,
					err,
					test.Output,
					test.Err,
				)
			}
		})
	}
}
