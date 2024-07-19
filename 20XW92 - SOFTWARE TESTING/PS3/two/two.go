package two

import "errors"

var ErrInvalidPrice = errors.New("Invalid price. Price can't be lesser than 0.")

func GetDiscountRate(price int) (int, error) {
	switch {
	case price < 1:
		{
			return 0, ErrInvalidPrice
		}
	case (price >= 1 && price <= 50):
		{
			return 0, nil
		}
	case (price > 50 && price <= 200):
		{
			return 5, nil
		}
	case (price > 200 && price <= 500):
		{
			return 10, nil
		}
	default:
		{
			return 15, nil
		}
	}
}
