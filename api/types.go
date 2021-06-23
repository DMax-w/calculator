package api

type response struct {
	Success bool    `json:"Success"`
	ErrCode string  `json:"ErrCode"`
	Value   float64 `json:"Value"`
}

type params struct {
	a float64
	b float64
}