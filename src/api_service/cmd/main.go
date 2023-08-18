package main

import (
	"flag"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/go-openapi/loads"
)

var port int
var host string

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&host, "host", "", "Host to listen on")
	flag.Parse()

	swaggerSpec, err := loads.Embedded(api.SwaggerJSON, api.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	adminAPI := operations.NewAdminAPI(swaggerSpec)
	server := api.NewServer(adminAPI)
	defer server.Shutdown()
	defer finalizeApp()

	handlers, err := initializeApp()
	if err != nil {
		log.Println("Initialization Error. Shutting down..")
		return
	}
	handlers.InstallHandlers(adminAPI)

	server.Port = port
	server.Host = host
	server.ConfigureAPI()

	log.Println("Starting Short URL API server..")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
