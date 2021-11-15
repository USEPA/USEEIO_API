package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
)

// Demand describes the final demand vector and the perspective for which a
// result should be calculated.
type Demand struct {
	Perspective string        `json:"perspective"`
	Entries     []DemandEntry `json:"demand"`
}

// DemandEntry describes a single entry in the final demand vector for the
// calculation.
type DemandEntry struct {
	SectorID string  `json:"sector"`
	Amount   float64 `json:"amount"`
}

// Result contains the result data of a calculation.
type Result struct {
	Indicators []string    `json:"indicators"`
	Sectors    []string    `json:"sectors"`
	Data       [][]float64 `json:"data"`
	Totals     []float64   `json:"totals"`
}

func (s *server) calculate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		modelDir := s.getModelDir(w, r)
		if modelDir == "" {
			return
		}

		decoder := json.NewDecoder(r.Body)
		var d Demand
		err := decoder.Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		result := calculate(modelDir, &d, w)
		if result != nil {
			serveJSON(result, w)
		}
	}
}

func calculate(dir string, d *Demand, w http.ResponseWriter) *Result {
	if d == nil {
		http.Error(w, "no demand", http.StatusBadRequest)
		return nil
	}

	// read the indicators and sectors
	indicators, err := readIndicators(dir)
	if err != nil {
		http.Error(w, "invalid model; no indicators", http.StatusBadRequest)
		return nil
	}
	sectors, err := readSectors(dir)
	if err != nil {
		http.Error(w, "invalid model; no sectors", http.StatusBadRequest)
		return nil
	}
	demand := demandVector(d, sectors, w)
	if demand == nil {
		return nil
	}

	// N is used for the total result; thus, always loaded
	N, err := loadMatrix(filepath.Join(dir, "N.bin"))
	if err != nil {
		http.Error(w, "could not load matrix N",
			http.StatusInternalServerError)
		return nil
	}

	// calculate the perspective result
	var data *Matrix
	switch d.Perspective {
	case "direct":
		L, err := loadMatrix(filepath.Join(dir, "L.bin"))
		if err == nil {
			D, err := loadMatrix(filepath.Join(dir, "D.bin"))
			if err == nil {
				s := L.ScaledColumnSums(demand)
				data = D.ScaleColumns(s)
			}
		}
	case "intermediate":
		L, err := loadMatrix(filepath.Join(dir, "L.bin"))
		if err == nil {
			s := L.ScaledColumnSums(demand)
			data = N.ScaleColumns(s)
		}
	case "final":
		data = N.ScaleColumns(demand)
	default:
		http.Error(w, "invalid perspective: "+d.Perspective,
			http.StatusBadRequest)
		return nil
	}

	if err != nil || data == nil {
		http.Error(w, "calculation failed", http.StatusInternalServerError)
		return nil
	}

	// finally, set the result data
	r := Result{}
	r.Totals = N.ScaledColumnSums(demand)
	if data != nil {
		r.Data = data.Slice2d()
	}
	r.Indicators = make([]string, len(indicators))
	for i := range indicators {
		r.Indicators[i] = indicators[i].ID
	}
	r.Sectors = make([]string, len(sectors))
	for i := range sectors {
		r.Sectors[i] = sectors[i].ID
	}
	return &r
}

func demandVector(d *Demand, sectors []*Sector, w http.ResponseWriter) []float64 {
	idx := make(map[string]int)
	for _, sector := range sectors {
		idx[sector.ID] = sector.Index
	}
	v := make([]float64, len(sectors))
	for _, e := range d.Entries {
		i, ok := idx[e.SectorID]
		if !ok {
			http.Error(w, "invalid sector ID "+e.SectorID,
				http.StatusBadRequest)
			return nil
		}
		v[i] = e.Amount
	}
	return v
}
