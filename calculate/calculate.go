package calculate

import "errors"

func Calculate(a, b float64, do string) (float64, error) {
	var result float64

	switch do {
	case Add:
		result = a + b
	case Sub:
		result = a - b
	case Mul:
		result = a * b
	case Div:
		if b == 0 {
			return 0, errors.New("division by 0")
		}

		result = a / b
	default:
		return 0, errors.New("undefined do")
	}

	return result, nil
}
