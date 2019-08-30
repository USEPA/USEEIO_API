package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
)

// WriteAccessOptions writes headers for a preflight request in CORS
// https://developer.mozilla.org/de/docs/Web/HTTP/Methods/OPTIONS
func WriteAccessOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers",
		"Content-Type, Access-Control-Allow-Headers")
}

// ServeJSON converts the given entity to a JSON string and writes it to the
// given response.
func ServeJSON(e interface{}, w http.ResponseWriter) {
	if e == nil {
		http.Error(w, "No data", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ServeJSONBytes(data, w)
}

// ServeJSONBytes writes the given data as JSON content to the given writer. It
// also sets the respective access control headers so that cross domain requests
// are supported.
func ServeJSONBytes(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	WriteAccessOptions(w)
	w.Write(data)
}

// ReadCSV reads all rows from the given CSV file.
func ReadCSV(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	return reader.ReadAll()
}
