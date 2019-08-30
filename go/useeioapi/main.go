package main

import (
	"encoding/json"
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

	log.Println("Create server with static files from:", args.StaticDir)

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(args.StaticDir))
	r.Handle("/", NoCache(fs))

	r.HandleFunc("/api/models", GetModels)
	r.HandleFunc("/api/{model}/sectors",
		HandleGetSectors(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/sectors/{id}",
		HandleGetSector(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators",
		HandleGetIndicators(args.DataDir)).Methods("GET")
	r.HandleFunc("/api/{model}/indicators/{id}",
		HandleGetIndicator(args.DataDir)).Methods("GET")

	r.HandleFunc("/api/", ModelDispatch)

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

// ServeJSON converts the given entity to a JSON string and writes it to the
// given response.
func ServeJSON(e interface{}, w http.ResponseWriter) {
	if e == nil {
		http.Error(w, "No data", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ServeJSONBytes(data, w)
}

// ServeJSONBytes writes the given data as JSON content to the given writer. It
// also sets the respective access control headers so that cross domain requests
// are supported.
func ServeJSONBytes(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	writeAccessOptions(w)
	w.Write(data)
}
