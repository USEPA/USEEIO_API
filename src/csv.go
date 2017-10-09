package main

import (
	"encoding/csv"
	"os"
)

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
