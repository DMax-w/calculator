package api

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	add = "add"
	sub = "sub"
	mul = "mul"
	div = "div"
)

func calculate(a, b float64, do string) (float64, error) {
	var result float64
	switch do {
	case add:
		result = a + b
	case sub:
		result = a - b
	case mul:
		result = a * b
	case div:
		result = a / b
	default:
		return 0, errors.New("undefined do")

	}
	return result, nil
}

func parseParams(r *http.Request) (*params, error) {
	a, err := strconv.ParseFloat(r.FormValue("a"), 64)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseFloat(r.FormValue("b"), 64)
	if err != nil {
		return nil, err
	}

	return &params{
		a: a,
		b: b,
	}, nil
}
