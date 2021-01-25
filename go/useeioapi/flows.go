package main

import (
	"errors"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// Flow describes an elementary flow an IO model.
type Flow struct {
	ID       string `json:"id"`
	Index    int    `json:"index"`
	Flowable string `json:"flowable"`
	Context  string `json:"context"`
	Unit     string `json:"unit"`
	UUID     string `json:"uuid"`
}

// ReadFlows reads the flows from the CSV file `flows.csv` in the data folder
// of the respective model. It returns them in a slice where the flows are
// sorted by their indices.
func ReadFlows(folder string) ([]*Flow, error) {
	path := filepath.Join(folder, "flows.csv")
	records, err := ReadCSV(path)
	if err != nil {
		return nil, err
	}

	flows := make([]*Flow, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		if len(row) < 5 {
			return nil, errors.New("error in " + path +
				": each row should have at least 5 columns")
		}
		flow := Flow{}
		if flow.Index, err = strconv.Atoi(row[0]); err != nil {
			return nil, err
		}
		flow.ID = row[1]
		flow.Flowable = row[2]
		flow.Context = row[3]
		flow.Unit = row[4]
		if len(row) > 5 {
			flow.UUID = row[5]
		}
		flows[flow.Index] = &flow
	}
	return flows, nil
}

// HandleGetFlows returns the handler for GET /api/{model}/flows
func HandleGetFlows(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		folder := filepath.Join(dataDir, model)
		flows, err := ReadFlows(folder)
		if err != nil {
			http.Error(w, "no flows for model "+model+" found",
				http.StatusNotFound)
			return
		}
		ServeJSON(flows, w)
	}
}

// HandleGetFlow returns the handler for GET /api/{model}/flows/{id}
func HandleGetFlow(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		model := mux.Vars(r)["model"]
		id := mux.Vars(r)["id"]
		folder := filepath.Join(dataDir, model)
		flows, err := ReadFlows(folder)
		if err != nil {
			http.Error(w, "no flows for model "+model+" found",
				http.StatusNotFound)
			return
		}
		for i := range flows {
			flow := flows[i]
			if flow.ID == id || flow.UUID == id {
				ServeJSON(flow, w)
				return
			}
		}
		http.Error(w, "no flow with id "+id+" for model "+model+" found",
			http.StatusNotFound)
	}
}
