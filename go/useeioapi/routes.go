package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (s *server) mountRoutes() {
	r := s.router

	r.HandleFunc("/api/models", s.getModels()).Methods("GET")
	r.HandleFunc("/api/{model}/demands", s.getDemands()).Methods("GET")
	r.HandleFunc("/api/{model}/demands/{id}", s.getDemand()).Methods("GET")
	r.HandleFunc("/api/{model}/sectors",
		HandleGetSectors(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/sectors/{id:.+}",
		HandleGetSector(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/flows",
		HandleGetFlows(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/flows/{id:.+}",
		HandleGetFlow(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators",
		HandleGetIndicators(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators/{id:.+}",
		HandleGetIndicator(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/matrix/{matrix}",
		HandleGetMatrix(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/calculate",
		HandleCalculate(dataDir)).Methods("POST")
	r.HandleFunc("/api/{model}/years",
		HandleGetYears(dataDir)).Methods("GET")

	// serve the crosswalk.csv file
	r.HandleFunc("/api/sectorcrosswalk.csv",
		func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(s.dataDir, "sectorcrosswalk.csv")
			if !fileExists(path) {
				http.Error(w, "sectorcrosswalk.csv does not exist",
					http.StatusNotFound)
				return
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				http.Error(w, "failed to read sectorcrosswalk.csv from data folder",
					http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/csv")
			WriteAccessOptions(w)
			if _, err := w.Write(data); err != nil {
				http.Error(w, "failed to serve sectorcrosswalk.csv from data folder",
					http.StatusInternalServerError)
			}
		}).Methods("GET")

	// handle CORS preflight requests
	r.PathPrefix("/api").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			WriteAccessOptions(w)
		}).Methods("OPTIONS")

	// check if we have static files to host
	if s.staticDir != "" {
		if stat, err := os.Stat(s.staticDir); err == nil && stat.IsDir() {
			log.Println("Host static files from:", s.staticDir)
			fs := http.FileServer(http.Dir(s.staticDir))
			r.PathPrefix("/").Handler(NoCache(fs))
		} else {
			log.Println("WARNING: ", s.staticDir,
				"is not a folder; will not host static files")
		}
	}
}
