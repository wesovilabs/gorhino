package main

import (
	"gihub.com/wesovilabs/gorhino/api"
	"gihub.com/wesovilabs/gorhino/logger"
	"gihub.com/wesovilabs/gorhino/props"
	"gihub.com/wesovilabs/gorhino/server"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("swat_demo_rest_api")

func main() {
	log.Info("Running application.")
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
