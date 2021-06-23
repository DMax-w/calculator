package api

import (
	"encoding/json"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var resp response

	p, err :=  parseParams(r)
	if err != nil {
		resp.ErrCode = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	result, err := calculate(p.a, p.b, add)
	if err != nil {
		resp.ErrCode = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Success = true
	resp.Value = result

	json.NewEncoder(w).Encode(resp)
}