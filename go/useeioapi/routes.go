package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// ModelDispatch reads the model and path (which should be /api/<model-id>/*)
// from the given request and dispatches to the respective function.
func ModelDispatch(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)
	if r.Method == "OPTIONS" {
		writeAccessOptions(w)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	model := models[parts[0]]
	if model == nil {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	cmd := parts[1]
	switch cmd {
	case "calculate":
		Calculate(model, w, r)
	}
}

// Calculate runs a calculation with the demand from the request body and
// returns the result.
func Calculate(model *Model, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var d Demand
	err := decoder.Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	calc, err := NewCalculator(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := calc.Calculate(&d)
	ServeJSON(result, w)
}
