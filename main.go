package main

import (
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/wesovilabs/taurus/api"
	"github.com/wesovilabs/taurus/logger"
	"github.com/wesovilabs/taurus/props"
	"github.com/wesovilabs/taurus/server"
)

var log = logging.MustGetLogger("swat_demo_rest_api")

func main() {
	log.Info("Running application.")
	api.InitializeBoltClient()
	properties := setUp()
	router := mux.NewRouter().StrictSlash(false)
	api.DefineDefaultHandlers(router)
	api.DefineRestAPI(router)
	server.ConfigureServerAndRun(router, properties)
}

func setUp() props.Properties {
	logger.InitializeLogger()
	return props.LoadProperties()
}
