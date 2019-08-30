package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Indicator describes an impact assessment indicator of the input-output model.
type Indicator struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Unit  string `json:"unit"`
	Group string `json:"group"`
}

// ReadIndicators reads the indicators from the CSV file in the data folder. It
// returns them in a slice where the indicators are sorted by their indices.
func ReadIndicators(folder string) ([]*Indicator, error) {
	path := filepath.Join(folder, "indicators.csv")
	records, err := ReadCSV(path)
	if err != nil {
		return nil, err
	}

	indicators := make([]*Indicator, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		if len(row) < 6 {
			return nil, errors.New("error in " + path +
				": each row should have 8 columns")
		}
		i := Indicator{}
		if i.Index, err = strconv.Atoi(row[0]); err != nil {
			return nil, err
		}
		i.ID = row[3]
		i.Name = row[2]
		i.Code = row[3]
		i.Unit = row[4]
		if i.Group, err = mapIndicatorGroup(row[5]); err != nil {
			return nil, err
		}
		indicators[i.Index] = &i
	}
	return indicators, nil
}

func mapIndicatorGroup(csvVal string) (string, error) {
	s := strings.ToLower(strings.TrimSpace(csvVal))
	switch s {
	case "impact potential":
		return "Impact Potential", nil
	case "resource use":
		return "Resource Use", nil
	case "waste generated":
		return "Waste Generated", nil
	case "economic & social":
		return "Economic & Social", nil
	case "chemical releases":
		return "Chemical Releases", nil
	default:
		return "", errors.New("Unknown indicator group: " + csvVal)
	}
}

// HandleGetIndicators returns the handler for GET /api/{model}/indicators
func HandleGetIndicators(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		folder := filepath.Join(dataDir, model)
		indicators, err := ReadIndicators(folder)
		if err != nil {
			http.Error(w, "no indicators for model "+model+" found",
				http.StatusNotFound)
			return
		}
		ServeJSON(indicators, w)
	}
}

// HandleGetIndicator returns the handler for GET /api/{model}/indicators/{id}
func HandleGetIndicator(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		id := mux.Vars(r)["id"]
		folder := filepath.Join(dataDir, model)
		indicators, err := ReadIndicators(folder)
		if err != nil {
			http.Error(w, "no indicators for model "+model+" found",
				http.StatusNotFound)
			return
		}
		for i := range indicators {
			indicator := indicators[i]
			if indicator.ID == id {
				ServeJSON(indicator, w)
				return
			}
		}
		http.Error(w, "no indicator with id "+id+" for model "+model+" found",
			http.StatusNotFound)
	}
}
