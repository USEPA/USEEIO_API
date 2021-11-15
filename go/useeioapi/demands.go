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
		d.ID = row[0]
		if d.Year, err = strconv.Atoi(row[1]); err != nil {
			return nil, err
		}
		d.Type = strings.TrimSpace(row[2])
		d.System = strings.TrimSpace(row[3])
		d.Location = strings.TrimSpace(row[4])
		demands[idx-1] = &d
	}
	return demands, nil
}

func (s *server) getDemands() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}
		demands, err := ReadDemandInfos(modelDir)
		if err != nil {
			http.Error(w, "failed to read demands", http.StatusInternalServerError)
			return
		}
		ServeJSON(demands, w)
	}
}

func (s *server) getDemand() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}
		id := mux.Vars(r)["id"]
		file := filepath.Join(modelDir, "demands", id+".json")
		data, err := ioutil.ReadFile(file)
		if err != nil {
			http.Error(w, "demand "+id+" not found", http.StatusNotFound)
			return
		}
		ServeJSONBytes(data, w)
	}
}
