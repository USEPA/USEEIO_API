package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

// A Model is an input-output model in the data folder.
type Model struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description,omitempty"`

	Folder      string        `json:"-"`
	Sectors     []*Sector     `json:"-"`
	Indicators  []*Indicator  `json:"-"`
	DemandInfos []*DemandInfo `json:"-"`

	sectorMap map[string]*Sector
	numCache  map[string]*Matrix
	dqiCache  map[string][][]string
}

// InitModels initializes the models from the given data folder.
func InitModels(dataDir string) map[string]*Model {
	models := make(map[string]*Model)
	rows, err := ReadCSV(filepath.Join(dataDir, "models.csv"))
	if err != nil {
		log.Fatal("failed to read models.csv", err)
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		model := &Model{
			ID:          row[0],
			Name:        row[1],
			Location:    row[2],
			Description: row[3]}
		model.Folder = filepath.Join(dataDir, model.ID)
		err = model.Init()
		if err != nil {
			log.Println("failed to load model in", model.Folder, err)
			continue
		}
		models[model.ID] = model
		log.Println("loaded model", model.Name)
	}
	return models
}

// Init initializes the model
func (m *Model) Init() error {
	sectors, err := ReadSectors(m.Folder)
	if err != nil {
		return err
	}
	sectorMap := make(map[string]*Sector)
	for i := range sectors {
		s := sectors[i]
		sectorMap[s.ID] = s
	}
	indicators, err := ReadIndicators(m.Folder)
	if err != nil {
		return err
	}
	demands, err := ReadDemandInfos(m.Folder)
	if err != nil {
		return err
	}
	m.Sectors = sectors
	m.Indicators = indicators
	m.DemandInfos = demands
	m.sectorMap = sectorMap
	m.numCache = make(map[string]*Matrix)
	m.dqiCache = make(map[string][][]string)
	return nil
}

// Sector returns the sector with the given ID.
func (m *Model) Sector(id string) *Sector {
	return m.sectorMap[id]
}

// SectorIDs returns the IDs of the sectors in the index order.
func (m *Model) SectorIDs() []string {
	ids := make([]string, len(m.Sectors))
	for i := range m.Sectors {
		ids[i] = m.Sectors[i].ID
	}
	return ids
}

// IndicatorIDs returns the IDs of the indicators in the index order.
func (m *Model) IndicatorIDs() []string {
	ids := make([]string, len(m.Indicators))
	for i := range m.Indicators {
		ids[i] = m.Indicators[i].ID
	}
	return ids
}

// Matrix returns the numeric matrix with the given name (e.g. `A`) from the
// model.
func (m *Model) Matrix(name string) (*Matrix, error) {
	matrix := m.numCache[name]
	if matrix != nil {
		return matrix, nil
	}
	file := filepath.Join(m.Folder, name+".bin")
	var err error
	matrix, err = Load(file)
	if err != nil {
		return nil, err
	}
	m.numCache[name] = matrix
	return matrix, nil
}

// DqiMatrix returns the DQI matrix with the given name (e.g. `B_dqi`) from the
// model.
func (m *Model) DqiMatrix(name string) ([][]string, error) {
	matrix := m.dqiCache[name]
	if matrix != nil {
		return matrix, nil
	}
	file := filepath.Join(m.Folder, name+".csv")
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	m.dqiCache[name] = records
	return records, nil
}
