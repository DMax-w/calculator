package api

import (
	"encoding/json"
	"fmt"
	"github.com/DMax-w/calculator/calculate"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	req  string
	resp response
}

func TestAddHandler(t *testing.T) {
	cases := []TestCase{
		{
			req:  getRequest(calculate.Add, 5, 5),
			resp: response{true, "", 10},
		},
	}
	for caseNum, item := range cases {
		req := httptest.NewRequest("GET", item.req, nil)
		w := httptest.NewRecorder()

		AddHandler(w, req)

		resp := w.Result()

		var bodyResp response

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&bodyResp); err != nil {
			t.Error("response type error")
			return
		}

		if bodyResp != item.resp {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyResp, item.resp)
		}
	}
}

func TestDivHandler(t *testing.T) {
	cases := []TestCase{
		{
			req:  getRequest(calculate.Div, 10, 5),
			resp: response{true, "", 2},
		},
		{
			req:  getRequest(calculate.Div, 10, 0),
			resp: response{false, "division by 0", 0},
		},
	}
	for caseNum, item := range cases {
		req := httptest.NewRequest("GET", item.req, nil)
		w := httptest.NewRecorder()

		DivHandler(w, req)

		resp := w.Result()

		var bodyResp response

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&bodyResp); err != nil {
			t.Error("response type error")
			return
		}

		if bodyResp != item.resp {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyResp, item.resp)
		}
	}
}

func TestMulHandler(t *testing.T) {
	cases := []TestCase{
		{
			req:  getRequest(calculate.Mul, 5, 5),
			resp: response{true, "", 25},
		},
		{
			req:  getRequest(calculate.Mul, 0, 5),
			resp: response{true, "", 0},
		},
	}
	for caseNum, item := range cases {
		req := httptest.NewRequest("GET", item.req, nil)
		w := httptest.NewRecorder()

		MulHandler(w, req)

		resp := w.Result()

		var bodyResp response

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&bodyResp); err != nil {
			t.Error("response type error")
			return
		}

		if bodyResp != item.resp {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyResp, item.resp)
		}
	}
}

func TestSubHandler(t *testing.T) {
	cases := []TestCase{
		{
			req:  getRequest(calculate.Add, 5, 5),
			resp: response{true, "", 0},
		},
	}
	for caseNum, item := range cases {
		req := httptest.NewRequest("GET", item.req, nil)
		w := httptest.NewRecorder()

		SubHandler(w, req)

		resp := w.Result()

		var bodyResp response

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&bodyResp); err != nil {
			t.Error("response type error")
			return
		}

		if bodyResp != item.resp {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseNum, bodyResp, item.resp)
		}
	}
}

func getRequest(do string, a, b int) string {
	return fmt.Sprintf("http://localhost:8080/api/%s?a=%v&b=%v", do, a, b)
}
