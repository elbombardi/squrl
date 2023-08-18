package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/util"
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
	adminAPI.Logger = func(s string, i ...interface{}) {
		util.Info(fmt.Sprintf(s, i...))
	}
	server := api.NewServer(adminAPI)
	defer server.Shutdown()
	defer finalizeApp()

	handlers, err := initializeApp()
	if err != nil {
		util.Error("Initialization Error. Shutting down..")
		return
	}
	handlers.InstallHandlers(adminAPI)

	server.Port = port
	server.Host = host
	server.ConfigureAPI()

	util.Info("Starting AdminAPI server..")
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
