package api

import (
	"encoding/json"
	"github.com/DMax-w/calculator/calculate"
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

	result, err := calculate.Calculate(p.a, p.b, calculate.Add)
	if err != nil {
		resp.ErrCode = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Success = true
	resp.Value = result

	json.NewEncoder(w).Encode(resp)
}