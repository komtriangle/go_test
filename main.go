package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/add", add).Methods("GET")
	router.HandleFunc("/api/sub", sub).Methods("GET")
	router.HandleFunc("/api/mul", mul).Methods("GET")
	router.HandleFunc("/api/div", div).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func PrepareParameters(r *http.Request) (int, int, error) {
	a, errA := ValidateParameter(r.URL.Query().Get("a"))
	b, errB := ValidateParameter(r.URL.Query().Get("b"))
	if errA != nil || errB != nil {
		errCode := ""
		if errA != nil {
			errCode += errA.Error()
		}
		if errB != nil {
			errCode += errB.Error()
		}
		return 0, 0, errors.New(errCode)
	} else {
		return a, b, nil
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	a, b, err := PrepareParameters(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resResponse := Response{Success: false, ErrCode: err.Error(), Value: 0}
		json.NewEncoder(w).Encode(&resResponse)
	} else {
		resResponse := Response{Success: true, ErrCode: "", Value: a + b}
		json.NewEncoder(w).Encode(&resResponse)
		w.WriteHeader(http.StatusOK)
	}
}
func sub(w http.ResponseWriter, r *http.Request) {
	a, b, err := PrepareParameters(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resResponse := Response{Success: false, ErrCode: err.Error(), Value: 0}
		json.NewEncoder(w).Encode(&resResponse)
	} else {
		resResponse := Response{Success: true, ErrCode: "", Value: a - b}
		json.NewEncoder(w).Encode(&resResponse)
		w.WriteHeader(http.StatusOK)
	}
}
func mul(w http.ResponseWriter, r *http.Request) {
	a, b, err := PrepareParameters(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resResponse := Response{Success: false, ErrCode: err.Error(), Value: 0}
		json.NewEncoder(w).Encode(&resResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		resResponse := Response{Success: true, ErrCode: "", Value: a * b}
		json.NewEncoder(w).Encode(&resResponse)

	}
}
func div(w http.ResponseWriter, r *http.Request) {
	a, b, err := PrepareParameters(r)
	if err == nil && b == 0 {
		err = errors.New("Divide by zero; ")
	}
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		resResponse := Response{Success: false, ErrCode: err.Error(), Value: 0}
		json.NewEncoder(w).Encode(&resResponse)
	} else {
		resResponse := Response{Success: true, ErrCode: "", Value: a / b}
		json.NewEncoder(w).Encode(&resResponse)
		w.WriteHeader(http.StatusOK)
	}
}
