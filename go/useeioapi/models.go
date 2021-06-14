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

// ReadModelInfos reads the model information from the data folder.
func ReadModelInfos(dataDir string) ([]*ModelInfo, error) {
	rows, err := ReadCSV(filepath.Join(dataDir, "models.csv"))
	if err != nil {
		log.Println("ERROR: failed to read models.csv", err)
		return nil, err
	}
	models := make([]*ModelInfo, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 || len(row) < 6 {
            log.Println("ERROR: Models.csv does not contain all required fields.")
            return nil
		}
		models = append(models, &ModelInfo{
			ID:           row[0],
			Name:         row[1],
			Location:     row[2],
			Description:  row[3],
			SectorSchema: row[4],
			Hash: row[5]})
	}
	return models, nil
}

// HandleGetModels returns the handler for GET /api/models
func HandleGetModels(dataDir string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		models, err := ReadModelInfos(dataDir)
		if err != nil {
			http.Error(w, "failed to read IO models",
				http.StatusInternalServerError)
			return
		}
		ServeJSON(models, w)
	}
}
