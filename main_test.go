package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Add(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/add", add)
	ts := httptest.NewServer(r)
	defer ts.Close()
	testData := []TestModel{
		TestModel{A: 5, B: 5, Response: Response{Success: true, ErrCode: "", Value: 10}},
		TestModel{A: 0, B: 0, Response: Response{Success: true, ErrCode: "", Value: 0}},
		TestModel{A: 1, B: 3, Response: Response{Success: true, ErrCode: "", Value: 4}},
		TestModel{A: 2, B: 5, Response: Response{Success: true, ErrCode: "", Value: 7}},
		TestModel{A: 100, B: 100, Response: Response{Success: true, ErrCode: "", Value: 200}},
	}
	for _, val := range testData {
		res, _ := http.Get(ts.URL + fmt.Sprintf("/api/add?a=%d&b=%d", val.A, val.B))
		var body Response
		json.NewDecoder(res.Body).Decode(&body)
		if body != val.Response {
			t.Errorf("Response is not correct")
		}
	}
}

func Test_Sub(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/sub", sub)
	ts := httptest.NewServer(r)
	defer ts.Close()
	testData := []TestModel{
		TestModel{A: 5, B: 5, Response: Response{Success: true, ErrCode: "", Value: 0}},
		TestModel{A: 0, B: 0, Response: Response{Success: true, ErrCode: "", Value: 0}},
		TestModel{A: 1, B: 3, Response: Response{Success: true, ErrCode: "", Value: -2}},
		TestModel{A: 2, B: 5, Response: Response{Success: true, ErrCode: "", Value: -3}},
		TestModel{A: 100, B: 100, Response: Response{Success: true, ErrCode: "", Value: 0}},
	}
	for _, val := range testData {
		res, _ := http.Get(ts.URL + fmt.Sprintf("/api/sub?a=%d&b=%d", val.A, val.B))
		var body Response
		json.NewDecoder(res.Body).Decode(&body)
		if body != val.Response {
			t.Errorf("Response is not correct")
		}
	}
}

func Test_Mul(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/mul", mul)
	ts := httptest.NewServer(r)
	defer ts.Close()
	testData := []TestModel{
		TestModel{A: 5, B: 5, Response: Response{Success: true, ErrCode: "", Value: 25}},
		TestModel{A: 0, B: 0, Response: Response{Success: true, ErrCode: "", Value: 0}},
		TestModel{A: 1, B: 3, Response: Response{Success: true, ErrCode: "", Value: 3}},
		TestModel{A: 2, B: 5, Response: Response{Success: true, ErrCode: "", Value: 10}},
		TestModel{A: 100, B: 100, Response: Response{Success: true, ErrCode: "", Value: 10000}},
	}
	for _, val := range testData {
		res, _ := http.Get(ts.URL + fmt.Sprintf("/api/mul?a=%d&b=%d", val.A, val.B))
		var body Response
		json.NewDecoder(res.Body).Decode(&body)
		if body != val.Response {
			t.Errorf("Response is not correct")
		}
	}
}

func Test_Div(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/div", div)
	ts := httptest.NewServer(r)
	defer ts.Close()
	testData := []TestModel{
		TestModel{A: 5, B: 5, Response: Response{Success: true, ErrCode: "", Value: 1}},
		TestModel{A: 0, B: 1, Response: Response{Success: true, ErrCode: "", Value: 0}},
		TestModel{A: 6, B: 3, Response: Response{Success: true, ErrCode: "", Value: 2}},
		TestModel{A: 50, B: 5, Response: Response{Success: true, ErrCode: "", Value: 10}},
		TestModel{A: 100, B: 100, Response: Response{Success: true, ErrCode: "", Value: 1}},
	}
	for _, val := range testData {
		res, _ := http.Get(ts.URL + fmt.Sprintf("/api/div?a=%d&b=%d", val.A, val.B))
		var body Response
		json.NewDecoder(res.Body).Decode(&body)
		if body != val.Response {
			t.Errorf("Response is not correct")
		}
	}
}
