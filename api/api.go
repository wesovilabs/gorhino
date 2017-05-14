package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const pathPrefix string = "/api/v0"

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
	// configure the router to always run this handler when it couldn't match a request to any other handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s not found\n", r.URL)))
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
