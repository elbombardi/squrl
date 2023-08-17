package main

import (
	"flag"
	"log"
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
		log.Println("Initialization Error. Shutting down..")
		return
	}
	log.Println("Starting Redirection server..")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
