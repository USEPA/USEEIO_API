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
	Index    int    `json:"index"`
	ID       string `json:"id"`
	Flowable string `json:"flowable"`
	Context  string `json:"context"`
	Unit     string `json:"unit"`
	UUID     string `json:"uuid"`
}

// readFlows reads the flows from the CSV file `flows.csv` in the data folder
// of the respective model. It returns them in a slice where the flows are
// sorted by their indices.
func readFlows(folder string) ([]*Flow, error) {
	path := filepath.Join(folder, "flows.csv")
	records, err := readCSV(path)
	if err != nil {
		return nil, err
	}

	flows := make([]*Flow, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}

		if len(row) < 6 {
			return nil, errors.New("error in " + path +
				": each row should have 6 columns")
		}
		flow := Flow{}
		if flow.Index, err = strconv.Atoi(row[0]); err != nil {
			return nil, err
		}
		flow.ID = row[1]
		flow.Flowable = row[2]
		flow.Context = row[3]
		flow.Unit = row[4]
		flow.UUID = row[5]
		flows[flow.Index] = &flow
	}
	return flows, nil
}

func (s *server) getFlows() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}

		flows, err := readFlows(modelDir)
		if err != nil {
			http.Error(w, "no flows found", http.StatusInternalServerError)
			return
		}
		serveJSON(flows, w)
	}
}

func (s *server) getFlow() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}

		id := mux.Vars(r)["id"]
		flows, err := readFlows(modelDir)
		if err != nil {
			http.Error(w, "no flows found", http.StatusInternalServerError)
			return
		}
		for i := range flows {
			flow := flows[i]
			if flow.ID == id || flow.UUID == id {
				serveJSON(flow, w)
				return
			}
		}
		http.Error(w, "no flow with id "+id+" found", http.StatusNotFound)
	}
}
