package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// DemandInfo contains the meta-data of a demand vector.
type DemandInfo struct {
	ID       string `json:"id"`
	Year     int    `json:"year"`
	Type     string `json:"type"`
	System   string `json:"system"`
	Location string `json:"location,omitempty"`
}

// ReadDemandInfos reads the meta-data of the available demand vectors from the
// CSV file in the data folder.
func ReadDemandInfos(folder string) ([]*DemandInfo, error) {
	path := filepath.Join(folder, "demands.csv")
	records, err := ReadCSV(path)
	if err != nil {
		return nil, err
	}

	demands := make([]*DemandInfo, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		if len(row) < 5 {
			return nil, errors.New("error in " + path +
				": each row should have 5 columns")
		}

		d := DemandInfo{}
		d.ID = row[0] // TODO: check file exists
		if d.Year, err = strconv.Atoi(row[1]); err != nil {
			return nil, err
		}
		if d.Type, err = mapDemandType(row[2]); err != nil {
			return nil, err
		}
		d.System = strings.TrimSpace(row[3])
		d.Location = strings.TrimSpace(row[4])
		demands[idx-1] = &d
	}
	return demands, nil
}

func mapDemandType(csvVal string) (string, error) {
	t := strings.ToLower(strings.TrimSpace(csvVal))
	switch t {
	case "production":
		return "Production", nil
	case "consumption":
		return "Consumption", nil
	default:
		return "", errors.New("Unknown demand type: " + csvVal)
	}
}

// HandleGetDemands returns the handler for GET /api/{model}/demands
func HandleGetDemands(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		folder := filepath.Join(dataDir, model)
		demands, err := ReadDemandInfos(folder)
		if err != nil {
			http.Error(w, "no demands for model "+model+" found",
				http.StatusNotFound)
			return
		}
		ServeJSON(demands, w)
	}
}

// HandleGetDemand returns the handler for GET /api/{model}/demands/{id}
func HandleGetDemand(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		id := mux.Vars(r)["id"]
		file := filepath.Join(dataDir, model, "demands", id+".json")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			http.Error(w, "no demand "+id+" for model "+model+" found",
				http.StatusNotFound)
			return
		}
		ServeJSONBytes(data, w)
	}
}
