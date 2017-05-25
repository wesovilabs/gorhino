package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swatlabs/taurus/db"
	"log"
	"net/http"
)

const pathPrefix string = "/api/v0"

var dBClient db.IBoltClient

//InitializeBoltClient Creates instance and calls the OpenBoltDb and Seed funcs
func InitializeBoltClient() {
	dBClient = &db.BoltClient{}
	dBClient.OpenBoltDb()
	dBClient.InitializeBucket()
}

//DefineRestAPI - Defining Rest API for application
func DefineRestAPI(router *mux.Router) {
	apiRouter := router.Headers("Content-Type", "application/json").PathPrefix(pathPrefix).Subrouter()
	if _, err := defineAnonymousResources(apiRouter); err != nil {
		fmt.Println(err)
		log.Fatal("Something failed while defining anonymous resources.")
	}
	var authenticatedRouter mux.Router
	if _, err := configureAuthenticatedRouter(apiRouter, &authenticatedRouter); err != nil {
		fmt.Println(err)
		log.Fatal("Something failed while defining anonymous resources.")
	}
	if _, err := defineAuthenticatedResources(&authenticatedRouter); err != nil {
		fmt.Println(err)
		log.Fatal("Something failed while defining authenticated resources.")
	}
}

//DefineDefaultHandlers - Defining Default Handlers for application
func DefineDefaultHandlers(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(fmt.Sprintf("%s not found\n", req.URL)))
	})

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func configureAuthenticatedRouter(router *mux.Router, authenticatedRouter *mux.Router) (*mux.Router, error) {
	authenticatedRouter = router.Headers("Content-Type", "application/json").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		adminToken := r.Header.Get("X-Authorization")
		if adminToken == "" {
			return false
		}
		return adminToken == "ValidToken"
	}).PathPrefix("/admin").Subrouter()
	return authenticatedRouter, nil
}

func defineAnonymousResources(router *mux.Router) (*mux.Router, error) {
	router.HandleFunc("/domains", createDomainHandler).Methods("POST").Name("CreateDomain")
	return router, nil
}

func defineAuthenticatedResources(router *mux.Router) (*mux.Router, error) {
	//router.HandleFunc("/{id:[0-9]+}", DefaultHandler).Methods("POST")
	return router, nil
}
