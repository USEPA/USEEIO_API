package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// Indicator describes an impact assessment indicator of the input-output model.
type Indicator struct {
	ID         string `json:"id"`
	Index      int    `json:"index"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Unit       string `json:"unit"`
	Group      string `json:"group"`
	SimpleUnit string `json:"simpleunit"`
	SimpleName string `json:"simplename"`
}

// readIndicators reads the indicators from the CSV file in the data folder. It
// returns them in a slice where the indicators are sorted by their indices.
func readIndicators(folder string) ([]*Indicator, error) {
	path := filepath.Join(folder, "indicators.csv")
	records, err := readCSV(path)
	if err != nil {
		return nil, err
	}

	indicators := make([]*Indicator, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		if len(row) < 8 {
			return nil, errors.New("error in " + path +
				": each row should have 8 columns")
		}
		i := Indicator{}
		if i.Index, err = strconv.Atoi(row[0]); err != nil {
			return nil, err
		}
		i.ID = row[1]
		i.Name = row[2]
		i.Code = row[3]
		i.Unit = row[4]
		i.Group = row[5]
		i.SimpleUnit = row[6]
		i.SimpleName = row[7]
		indicators[i.Index] = &i
	}
	return indicators, nil
}

func (s *server) getIndicators() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}
		indicators, err := readIndicators(modelDir)
		if err != nil {
			http.Error(w, "no indicators found", http.StatusInternalServerError)
			return
		}
		serveJSON(indicators, w)
	}
}

func (s *server) getIndicator() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}
		id := mux.Vars(r)["id"]

		indicators, err := readIndicators(modelDir)
		if err != nil {
			http.Error(w, "no indicators found", http.StatusInternalServerError)
			return
		}
		for i := range indicators {
			indicator := indicators[i]
			if indicator.ID == id {
				serveJSON(indicator, w)
				return
			}
		}
		http.Error(w, "no indicator with id "+id+" found", http.StatusNotFound)
	}
}
