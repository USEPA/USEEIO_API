package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// ModelInfo describes an input-output model in the data folder.
type ModelInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Description  string `json:"description,omitempty"`
	SectorSchema string `json:"sectorschema,omitempty"`
	Hash         string `json:"hash,omitempty"`
}

// readModelInfos reads the model information from the data folder.
func readModelInfos(dataDir string) ([]*ModelInfo, error) {
	rows, err := readCSV(filepath.Join(dataDir, "models.csv"))
	if err != nil {
		log.Println("ERROR: failed to read models.csv", err)
		return nil, err
	}
	models := make([]*ModelInfo, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue // skip header row
		}
		if len(row) < 6 {
			log.Println("ERROR: invalid models.csv: row ",
				i, " has less than 6 columns")
			return nil, err
		}
		models = append(models, &ModelInfo{
			ID:           row[0],
			Name:         row[1],
			Location:     row[2],
			Description:  row[3],
			SectorSchema: row[4],
			Hash:         row[5],
		})
	}
	return models, nil
}

func (s *server) getModels() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		serveJSON(s.models, w)
	}
}
