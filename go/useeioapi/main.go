package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var models map[string]*Model

func main() {
	args := GetArgs()

	log.Println("load data from folder:", args.DataDir)
	models = InitModels(args.DataDir)

	r := mux.NewRouter()

	log.Println("Create server with static files from:", args.StaticDir)
	fs := http.FileServer(http.Dir(args.StaticDir))
	r.Handle("/", NoCache(fs))

	// model routes
	r.HandleFunc("/api/models", GetModels)
	r.HandleFunc("/api/{model}/demands",
		HandleGetDemands(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/demands/{id}",
		HandleGetDemand(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/sectors",
		HandleGetSectors(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/sectors/{id}",
		HandleGetSector(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators",
		HandleGetIndicators(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators/{id}",
		HandleGetIndicator(args.DataDir)).Methods("GET")

	// matrix routes
	r.HandleFunc("/api/{model}/matrix/{matrix}",
		HandleGetMatrix(args.DataDir)).Methods("GET")

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

// GetModels returns the a of models.
func GetModels(w http.ResponseWriter, r *http.Request) {
	list := make([]*Model, len(models))
	i := 0
	for key := range models {
		list[i] = models[key]
		i++
	}
	ServeJSON(list, w)
}
