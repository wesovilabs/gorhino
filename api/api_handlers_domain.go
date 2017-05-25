package api

import (
	"encoding/json"
	"github.com/swatlabs/taurus/checker"
	"github.com/swatlabs/taurus/model"
	"log"
	"net/http"
)

func createDomainHandler(res http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}
	decoder := json.NewDecoder(req.Body)
	var domain model.Domain
	if err := decoder.Decode(&domain); err != nil {
		log.Fatal("Error while decoding")
		respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}
	defer req.Body.Close()
	if err := domain.IsValid(); err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}
	currentDomain, err := dBClient.QueryDomain(domain.UID)
	if err != nil {

	}
	if currentDomain.UID == domain.UID {
		respondWithError(res, http.StatusConflict, domain.UID+"DomainUID already in use")
		checker.RegisterDomain(currentDomain)
		return
	}
	dBClient.CreateDomain(domain)
	respondWithJSON(res, http.StatusCreated, domain)
}
