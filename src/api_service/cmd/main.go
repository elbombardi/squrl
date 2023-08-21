package main

import (
	"flag"
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
		slog.Error("Initialization Error. Shutting down", "Details", err)
		return
	}
	defer server.Shutdown()

	server.Port = port
	server.Host = host
	server.ConfigureAPI()

	slog.Info("Starting AdminAPI server..")
	err = server.Serve()
	if err != nil {
		slog.Error("Unexpected error", "Details", err)
	}

}
