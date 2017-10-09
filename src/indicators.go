package main

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
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
