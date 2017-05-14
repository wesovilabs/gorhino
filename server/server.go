package server

import (
	"github.com/op/go-logging"
	"github.com/wesovilabs/taurus/props"
	"net/http"
	"time"
)

var log = logging.MustGetLogger("swat_demo_rest_api/server")

//ConfigureServerAndRun - configuring serer and running it
func ConfigureServerAndRun(router http.Handler, properties props.Properties) {
	server := configureServer(router, properties)
	log.Infof("Launching server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func configureServer(router http.Handler, properties props.Properties) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         properties.Server.ServerAddress(),
		WriteTimeout: properties.Server.WriteTimeout * time.Second,
		ReadTimeout:  properties.Server.ReadTimeout * time.Second,
	}
}
