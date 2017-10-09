package main

import (
	"log"
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

// Calculator holds the data for a calculation.
type Calculator struct {
	model *Model

	D *Matrix
	L *Matrix
	U *Matrix
}

// Result contains the result data of a calculation.
type Result struct {
	Indicators []string    `json:"indicators"`
	Sectors    []string    `json:"sectors"`
	Data       [][]float64 `json:"data"`
	Totals     []float64   `json:"totals"`
}

// NewCalculator creates a new calculator from the given model.
func NewCalculator(model *Model) (*Calculator, error) {
	c := Calculator{model: model}
	var err error
	c.D, err = model.Matrix("D")
	if err != nil {
		return nil, err
	}
	c.L, err = model.Matrix("L")
	if err != nil {
		return nil, err
	}
	c.U, err = model.Matrix("U")
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Calculate calculates the given demand and returns the respective result.
func (c *Calculator) Calculate(d *Demand) *Result {
	if c == nil || d == nil {
		return nil
	}

	r := Result{}
	r.Indicators = c.model.IndicatorIDs()
	r.Sectors = c.model.SectorIDs()

	demand := c.demandVector(d)
	r.Totals = c.U.ScaledColumnSums(demand)

	var data *Matrix
	switch d.Perspective {
	case "direct":
		s := c.L.ScaledColumnSums(demand)
		data = c.D.ScaleColumns(s)
	case "intermediate":
		s := c.L.ScaledColumnSums(demand)
		data = c.U.ScaleColumns(s)
	case "final":
		data = c.U.ScaleColumns(demand)
	default:
		// TODO: log error
		return nil
	}

	r.Data = make([][]float64, data.Rows)
	for row := 0; row < data.Rows; row++ {
		rowVals := make([]float64, data.Cols)
		r.Data[row] = rowVals
		for col := 0; col < data.Cols; col++ {
			rowVals[col] = data.Get(row, col)
		}
	}
	return &r
}

func (c *Calculator) demandVector(d *Demand) []float64 {
	v := make([]float64, c.L.Rows)
	for _, e := range d.Entries {
		sector := c.model.Sector(e.SectorID)
		if sector == nil {
			log.Println("failed to get sector", e.SectorID, "for calculation")
			continue
		}
		v[sector.Index] = e.Amount
	}
	return v
}
