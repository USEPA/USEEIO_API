package main

import (
	"os"
	"strings"
)

// Args contains the command line arguments of the server application. All
// arguments are optional and should have reasonable default values.
type Args struct {

	// A path to a folder that contains matrices and meta-data files.
	DataDir string

	// A path to a folder that contains the static resources (HTML, JavaScript,
	// CSS) of the application.
	StaticDir string

	// The port that is used for the HTTP server.
	Port string
}

// GetArgs returns the programm arguments of the server application.
func GetArgs() *Args {
	args := Args{
		DataDir:   "data",
		StaticDir: "static",
		Port:      "80"}
	for flag, arg := range readArgs() {
		switch flag {
		case "-data":
			args.DataDir = arg
		case "-static":
			args.StaticDir = arg
		case "-port":
			args.Port = arg
		}
	}
	// If the server runs in Cloud Foundry we need
	// to bind the port to the PORT environment variable
	port := os.Getenv("PORT")
	if port != "" {
		args.Port = port
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
