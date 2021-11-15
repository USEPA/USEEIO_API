package main

import (
	"encoding/csv"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
)

func (s *server) getMatrix() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}

		name := mux.Vars(r)["matrix"]
		if !isValidMatrixName(name) {
			http.Error(w, "invalid matrix name", http.StatusBadRequest)
			return
		}

		// get possible row or column
		col, err := indexParam("col", r.URL, w)
		if err != nil {
			return
		}
		row, err := indexParam("row", r.URL, w)
		if err != nil {
			return
		}

		// if the matrix ends with _dqi we serve it as DQI matrix
		// from a CSV file
		if strings.HasSuffix(name, "_dqi") {
			file := filepath.Join(modelDir, name+".csv")
			if !fileExists(file) {
				http.Error(w, "DQI matrix "+name+" does not exist",
					http.StatusNotFound)
				return
			}
			dqis, err := readDqiMatrix(file)
			if err != nil {
				http.Error(w, "Failed to load DQI matrix "+name,
					http.StatusInternalServerError)
				return
			}
			serveDqiMatrix(dqis, row, col, w)
			return
		}

		// otherwise we try to load it from our binary format
		file := filepath.Join(modelDir, name+".bin")
		if !fileExists(file) {
			http.Error(w, "Matrix "+name+" does not exist",
				http.StatusNotFound)
			return
		}
		matrix, err := LoadMatrix(file)
		if err != nil {
			http.Error(w, "Failed to load matrix "+name,
				http.StatusInternalServerError)
			return
		}
		serveMatrix(matrix, row, col, w)

	}
}

func serveMatrix(matrix *Matrix, row int, col int, w http.ResponseWriter) {
	// return a single column
	if col > -1 {
		if col >= matrix.Cols {
			http.Error(w, "Column out of bounds", http.StatusBadRequest)
			return
		}
		ServeJSON(matrix.Col(col), w)
		return
	}

	// return a single row
	if row > -1 {
		if row >= matrix.Rows {
			http.Error(w, "Row out of bound", http.StatusBadRequest)
			return
		}
		ServeJSON(matrix.Row(row), w)
		return
	}

	// return the full matrix
	ServeJSON(matrix.Slice2d(), w)
}

func readDqiMatrix(file string) ([][]string, error) {
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

func serveDqiMatrix(dqis [][]string, row int, col int, w http.ResponseWriter) {
	// return a single column
	if col > -1 {
		vals := make([]string, len(dqis))
		for row, rowVals := range dqis {
			if rowVals == nil || len(rowVals) <= col {
				http.Error(w, "Column out of bounds", http.StatusBadRequest)
				return
			}
			vals[row] = rowVals[col]
		}
		ServeJSON(vals, w)
		return
	}

	// return a single row
	if row > -1 {
		if row >= len(dqis) {
			http.Error(w, "Row out of bounds", http.StatusBadRequest)
			return
		}
		ServeJSON(dqis[row], w)
		return
	}

	// return the full matrix
	ServeJSON(dqis, w)
}

func indexParam(name string, reqURL *url.URL, w http.ResponseWriter) (int, error) {
	str := reqURL.Query().Get(name)
	if str == "" {
		return -1, nil
	}
	idx, err := strconv.Atoi(str)
	if err != nil || idx < 0 {
		http.Error(w, "Invalid index: "+name+"="+str, http.StatusBadRequest)
		return -1, err
	}
	return idx, nil
}

// Checks if the given name is a valid matrix name. A matrix name must start
// with a letter and can only consist of letters, digits, and underscores.
func isValidMatrixName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for i, char := range name {
		if i == 0 && !unicode.IsLetter(char) {
			return false
		}
		if unicode.IsLetter(char) ||
			unicode.IsDigit(char) ||
			char == '_' {
			continue
		}
		return false
	}
	return true
}
