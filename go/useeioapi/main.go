package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	args := GetArgs()

	// check that the data folder exists
	dataDir := args.DataDir
	if stat, err := os.Stat(dataDir); err != nil || !stat.IsDir() {
		log.Println("ERROR:", dataDir, "is not a valid data folder")
		os.Exit(1)
	}
	log.Println("Load data from folder:", dataDir)

	r := mux.NewRouter()

	// mount API routes
	r.HandleFunc("/api/models",
		HandleGetModels(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/demands",
		HandleGetDemands(dataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/demands/{id}",
		HandleGetDemand(dataDir)).Methods("GET")
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
			path := filepath.Join(dataDir, "sectorcrosswalk.csv")
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
	if args.StaticDir != "" {
		if stat, err := os.Stat(args.StaticDir); err == nil && stat.IsDir() {
			log.Println("Host static files from:", args.StaticDir)
			fs := http.FileServer(http.Dir(args.StaticDir))
			r.PathPrefix("/").Handler(NoCache(fs))
		} else {
			log.Println("WARNING: ", args.StaticDir,
				"is not a folder; will not host static files")
		}
	}

	// register shutdown hook
	log.Println("Register shutdown routines")
	ossignals := make(chan os.Signal)
	signal.Notify(ossignals, syscall.SIGTERM)
	signal.Notify(ossignals, syscall.SIGINT)
	go func() {
		<-ossignals
		log.Println("Shutdown server")
		os.Exit(0)
	}()

	log.Println("Starting server at port:", args.Port)
	http.ListenAndServe(":"+args.Port, r)
}
