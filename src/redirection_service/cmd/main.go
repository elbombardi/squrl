package main

import (
	"flag"
	"log"
	"log/slog"
)

var port int
var host string

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&host, "host", "", "Host to listen on")
	flag.Parse()
	defer finalizeApp()

	server, err := initializeApp()
	if err != nil {
		slog.Error("Initialization Error. Shutting down..")
		return
	}
	slog.Info("Starting Redirection server..")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
