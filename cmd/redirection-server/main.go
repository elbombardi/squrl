package main

import (
	"flag"
	"log"

	"github.com/elbombardi/squrl/redirection_service"
)

var port int
var host string

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&host, "host", "0.0.0.0", "Host to listen on")
	flag.Parse()

	routes, err := initializeApp()
	if err != nil {
		log.Println("Initialization Error. Shutting down..")
		return
	}
	server := redirection_service.NewServer(port, host, routes)

	log.Println("Starting Redirection server..")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
