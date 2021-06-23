package api

import (
	"net/http"
	"strconv"
)

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
