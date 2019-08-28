package main

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
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
