package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var models map[string]*Model

func main() {
	args := GetArgs()

	log.Println("load data from folder:", args.DataDir)
	initModels(args.DataDir)

	log.Println("Create server with static files from:", args.StaticDir)
	h := http.NewServeMux()
	fs := http.FileServer(http.Dir(args.StaticDir))
	h.Handle("/", NoCache(fs))
	h.HandleFunc("/api/models", GetModels)
	h.HandleFunc("/api/", ModelDispatch)

	log.Println("Starting server at port:", args.Port)
	http.ListenAndServe(":"+args.Port, h)
}

func initModels(dataDir string) {
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal("failed to read data dir.:", err)
	}
	models = make(map[string]*Model)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		fName := f.Name()
		path := filepath.Join(dataDir, fName)
		m, err := NewModel(fName, path)
		if err != nil {
			log.Println("failed to load model in", fName, err)
			continue
		}
		models[m.ID] = m
		log.Println("loaded model", m.Name)
	}
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
