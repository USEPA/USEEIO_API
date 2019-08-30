package main

import (
	"encoding/csv"
	"os"
)

// ReadDqiMatrix reads a DQI matrix from the given file.
func ReadDqiMatrix(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
