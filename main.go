package main

import (
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/swatlabs/taurus/api"
	"github.com/swatlabs/taurus/configuration"
	"github.com/swatlabs/taurus/server"
)

var log = logging.MustGetLogger("taurus")

func main() {
	log.Info("Running application.")
	api.InitializeBoltClient()
	properties := setUp()
	router := mux.NewRouter().StrictSlash(false)
	api.DefineDefaultHandlers(router)
	api.DefineRestAPI(router)
	server.ConfigureServerAndRun(router, properties)
}

func setUp() configuration.Properties {
	configuration.InitializeLogger()
	return configuration.LoadProperties()
}
