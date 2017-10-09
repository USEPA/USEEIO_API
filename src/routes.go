package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
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
	case "sectors":
		ServeJSON(model.Sectors, w)
	case "indicators":
		ServeJSON(model.Indicators, w)
	case "demands":
		GetDemands(model, parts, w)
	case "matrix":
		GetMatrix(model, w, r)
	case "calculate":
		Calculate(model, w, r)
	}
}

// Write headers for a preflight request in CORS
// https://developer.mozilla.org/de/docs/Web/HTTP/Methods/OPTIONS
func writeAccessOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, Access-Control-Allow-Headers")
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

// GetDemands returns list with all available demand vectors of the model or a
// specific demand vector with the given ID if the third path parameter is given.
func GetDemands(model *Model, path []string, w http.ResponseWriter) {
	var file string
	if len(path) < 3 || path[2] == "" {
		ServeJSON(model.DemandInfos, w)
		return
	}
	file = filepath.Join(model.Folder, "demands", path[2]+".json")
	data, err := ioutil.ReadFile(file)
	if err != nil {
		http.Error(w, file+" not found", http.StatusNotFound)
		return
	}
	ServeJSONBytes(data, w)
}
