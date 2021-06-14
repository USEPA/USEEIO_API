package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// Year describes a year.
type Year struct {
	ID          string `json:"id"`
	Index       int    `json:"index"`
}

// ReadYears reads the years from the CSV file in the data folder. It
// returns them in a slice where the years are sorted by their indices.
func ReadYears(folder string) ([]*Year, error) {
	path := filepath.Join(folder, "years.csv")
	records, err := ReadCSV(path)
	if err != nil {
		return nil, err
	}

	years := make([]*Year, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		y := Year{}
		y.Index, err = strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		y.ID = row[1]
		years[y.Index] = &y
	}
	return years, nil
}

// HandleGetSectors returns the handler for GET /api/{model}/years
func HandleGetYears (dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		folder := filepath.Join(dataDir, model)
		years, err := ReadYears(folder)
		if err != nil {
			http.Error(w, "no years for model "+model+" found",
				http.StatusNotFound)
			return
		}
		ServeJSON(years, w)
	}
}
