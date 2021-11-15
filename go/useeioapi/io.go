package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
)

// writeAccessOptions writes headers for a preflight request in CORS
// https://developer.mozilla.org/de/docs/Web/HTTP/Methods/OPTIONS
func writeAccessOptions(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}

// serveJSON converts the given entity to a JSON string and writes it to the
// given response.
func serveJSON(e interface{}, w http.ResponseWriter) {
	if e == nil {
		http.Error(w, "No data", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	serveJSONBytes(data, w)
}

// serveJSONBytes writes the given data as JSON content to the given writer. It
// also sets the respective access control headers so that cross domain requests
// are supported.
func serveJSONBytes(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	writeAccessOptions(w)
	w.Write(data)
}

// readCSV reads all rows from the given CSV file.
func readCSV(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	return reader.ReadAll()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
