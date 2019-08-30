package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	r.HandleFunc("/api/{model}/sectors", HandleGetSectors(args.DataDir))
	r.HandleFunc("/api/", ModelDispatch)

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
