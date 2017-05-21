package main

import (
	"net/http"
	"time"
)


func configureServerAndRun(router http.Handler, properties Properties) {
	server := configureServer(router, properties)
	log.Infof("Launching server on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func configureServer(router http.Handler, properties Properties) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         properties.Server.ServerAddress(),
		WriteTimeout: properties.Server.WriteTimeout * time.Second,
		ReadTimeout:  properties.Server.ReadTimeout * time.Second,
	}
}
