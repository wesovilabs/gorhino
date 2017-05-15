package api

import (
	"encoding/json"
	"github.com/wesovilabs/taurus/model"
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
	if err := domain.isValid(); err != nil {
		respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}
	/**
	if err := json.MarshalJSON(domain); err != nil {
		panic(err)
		respondWithError(res,http.StatusInternalServerError,err.Error())
		return
	}
	**/
	respondWithJSON(res, http.StatusCreated, domain)
}
