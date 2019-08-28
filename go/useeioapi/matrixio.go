package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetMatrix returns a matrix as JSON object.
func GetMatrix(model *Model, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/matrix/")
	if len(parts) < 2 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
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

	name := parts[1]
	switch name {
	case "A", "B", "C", "D", "L", "U":
		matrix, err := model.Matrix(name)
		if err != nil {
			http.Error(w, "Failed to load matrix", http.StatusInternalServerError)
			return
		}
		ServeMatrix(matrix, row, col, w)
	case "B_dqi", "D_dqi", "U_dqi":
		dqis, err := model.DqiMatrix(name)
		if err != nil {
			http.Error(w, "Failed to load matrix", http.StatusInternalServerError)
			return
		}
		ServeDqiMatrix(dqis, row, col, w)
	default:
		http.Error(w, "Unknown matrix: "+name, http.StatusNotFound)
	}
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
	return idx, err
}

// ServeMatrix serves the given numeric matrix as JSON object.
func ServeMatrix(matrix *Matrix, row int, col int, w http.ResponseWriter) {
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

// ServeDqiMatrix serves the given DQI matrix as JSON object.
func ServeDqiMatrix(dqis [][]string, row int, col int, w http.ResponseWriter) {
	// return a single column
	if col > -1 {
		serveDqiColumn(dqis, col, w)
		return
	}

	// return a single row
	if row > -1 {
		if row >= len(dqis) {
			http.Error(w, "Row out of bound", http.StatusBadRequest)
			return
		}
		ServeJSON(dqis[row], w)
		return
	}

	// return the full matrix
	ServeJSON(dqis, w)
}

func serveDqiColumn(dqis [][]string, col int, w http.ResponseWriter) {
	vals := make([]string, len(dqis))
	for row, rowVals := range dqis {
		if rowVals == nil || len(rowVals) <= col {
			http.Error(w, "Column out of bounds", http.StatusBadRequest)
			return
		}
		vals[row] = rowVals[col]
	}
	ServeJSON(vals, w)
}
