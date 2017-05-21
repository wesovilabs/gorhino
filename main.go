package main

import (
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"github.com/wesovilabs/taurus/api"
)

var log = logging.MustGetLogger("swat_demo_rest_api")

func main() {
	log.Info("Running application.")
	api.InitializeBoltClient()
	properties := setUp()
	router := mux.NewRouter().StrictSlash(false)
	api.DefineDefaultHandlers(router)
	api.DefineRestAPI(router)
	configureServerAndRun(router, properties)
}

func setUp() Properties {
	InitializeLogger()
	return LoadProperties()
}
