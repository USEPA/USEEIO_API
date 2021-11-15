package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	args := GetArgs()

	// check that the data folder exists
	server, err := newServer(args)
	if err != nil {
		log.Println("failed to create server", err)
		return
	}

	server.mountRoutes()

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

	server.start()
}
