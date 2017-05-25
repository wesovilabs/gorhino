package server

import (
	"github.com/op/go-logging"
	"github.com/swatlabs/taurus/configuration"
	"net/http"
	"time"
)

var log = logging.MustGetLogger("taurus_server")

//ConfigureServerAndRun run server
func ConfigureServerAndRun(router http.Handler, properties configuration.Properties) {
	server := configureServer(router, properties)
	log.Infof("Launching server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func configureServer(router http.Handler, properties configuration.Properties) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         properties.Server.ServerAddress(),
		WriteTimeout: properties.Server.WriteTimeout * time.Second,
		ReadTimeout:  properties.Server.ReadTimeout * time.Second,
	}
}
