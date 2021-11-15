package main

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// Sector describes an industry sector in an input-output model.
type Sector struct {
	ID          string `json:"id"`
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

// readSectors reads the sectors from the CSV file in the data folder. It
// returns them in a slice where the sectors are sorted by their indices.
func readSectors(folder string) ([]*Sector, error) {
	path := filepath.Join(folder, "sectors.csv")
	records, err := readCSV(path)
	if err != nil {
		return nil, err
	}

	sectors := make([]*Sector, len(records)-1)
	for idx, row := range records {
		if idx == 0 {
			continue
		}
		s := Sector{}
		s.Index, err = strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		s.ID = row[1]
		s.Name = row[2]
		s.Code = row[3]
		s.Location = row[4]
		s.Description = row[5]
		sectors[s.Index] = &s
	}
	return sectors, nil
}

func (s *server) getSectors() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}
		sectors, err := readSectors(modelDir)
		if err != nil {
			http.Error(w, "no sectors found", http.StatusInternalServerError)
			return
		}
		serveJSON(sectors, w)
	}
}

func (s *server) getSector() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}

		id := mux.Vars(r)["id"]
		sectors, err := readSectors(modelDir)
		if err != nil {
			http.Error(w, "no sectors found", http.StatusInternalServerError)
			return
		}
		for i := range sectors {
			s := sectors[i]
			if s.ID == id {
				serveJSON(s, w)
				return
			}
		}
		http.Error(w, "no sector with id "+id+" found", http.StatusNotFound)
	}
}
