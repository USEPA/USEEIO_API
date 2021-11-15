package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type server struct {
	dataDir   string
	staticDir string
	port      string
	models    []*ModelInfo
	router    *mux.Router
}

func newServer(args *args) (*server, error) {
	if stat, err := os.Stat(args.dataDir); err != nil || !stat.IsDir() {
		return nil, errors.New("invalid data folder: " + args.dataDir)
	}
	log.Println("Load data from folder:", args.dataDir)

	models, err := readModelInfos(args.dataDir)
	if err != nil {
		return nil, err
	}

	// check that there is a folder for each model
	for _, model := range models {
		modelDir := filepath.Join(args.dataDir, model.ID)
		stat, err := os.Stat(modelDir)
		if err != nil || !stat.IsDir() {
			return nil, errors.New("no model folder found for: " + model.ID)
		}
	}

	server := &server{
		dataDir:   args.dataDir,
		staticDir: args.staticDir,
		port:      args.port,
		models:    models,
		router:    mux.NewRouter()}
	return server, nil
}

func (s *server) start() {
	log.Println("Starting server at port:", s.port)
	http.ListenAndServe(":"+s.port, s.router)
}

func (s *server) isModel(id string) bool {
	for _, info := range s.models {
		if info.ID == id {
			return true
		}
	}
	return false
}

func (s *server) getModelDir(w http.ResponseWriter, r *http.Request) string {
	model := mux.Vars(r)["model"]
	if !s.isModel(model) {
		http.Error(w, "'"+model+"' is not a model", http.StatusBadRequest)
		return ""
	}
	return filepath.Join(s.dataDir, model)
}
