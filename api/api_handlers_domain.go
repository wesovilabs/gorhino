package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type domainSettings struct {
	domainType string
	public     bool
	sandbox    bool
}

//
type domain struct {
	uid      string
	name     string
	settings domainSettings
}

type credentials struct {
	username string
	password string
}

type gitConfiguration struct {
	credentials credentials
	uri         string
	branch      string
}

type fileSystemConfiguration struct {
	credentials credentials
	path        string
}

func (domain domain) isValid() error {
	fmt.Print("Result:")
	fmt.Print(domain.uid)
	if domain.uid == "" {
		return errors.New("Missing uuid")
	}
	return nil
}

func createDomainHandler(res http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request body")
		return
	}
	decoder := json.NewDecoder(req.Body)
	var domain domain
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
