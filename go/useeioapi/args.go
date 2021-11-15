package main

import (
	"os"
	"strings"
)

// args contains the command line arguments of the server application. All
// arguments are optional and should have reasonable default values.
type args struct {

	// A path to a folder that contains matrices and meta-data files.
	dataDir string

	// A path to a folder that contains the static resources (HTML, JavaScript,
	// CSS) of the application.
	staticDir string

	// The port that is used for the HTTP server.
	port string
}

// GetArgs returns the programm arguments of the server application.
func GetArgs() *args {
	args := args{
		dataDir:   "data",
		staticDir: "static",
		port:      "80"}
	for flag, arg := range readArgs() {
		switch flag {
		case "-data":
			args.dataDir = arg
		case "-static":
			args.staticDir = arg
		case "-port":
			args.port = arg
		}
	}
	// If the server runs in Cloud Foundry we need
	// to bind the port to the PORT environment variable
	port := os.Getenv("PORT")
	if port != "" {
		args.port = port
	}
	return &args
}

func readArgs() map[string]string {
	args := make(map[string]string)
	if len(os.Args) < 2 {
		return args
	}
	flag := ""
	for i, val := range os.Args {
		if i == 0 {
			continue
		}
		arg := strings.TrimSpace(val)
		if flag != "" {
			args[flag] = arg
			flag = ""
			continue
		}
		if strings.HasPrefix(arg, "-") {
			flag = arg
		}
	}
	return args
}
