package model

import (
	"errors"
	"fmt"
)

//DomainSettings structure
type DomainSettings struct {
	domainType string
	public     bool
	sandbox    bool
}

//Domain structure
type Domain struct {
	UID      string
	name     string
	settings DomainSettings
}

//Credentials structure
type Credentials struct {
	username string
	password string
}

//GitConfiguration structure
type GitConfiguration struct {
	credentials Credentials
	uri         string
	branch      string
}

//FileSystemConfiguration structure
type FileSystemConfiguration struct {
	credentials Credentials
	path        string
}

//IsValid Determines if a domain is valid
func (domain Domain) IsValid() error {
	fmt.Print("Result:")
	fmt.Print(domain.UID)
	if domain.UID == "" {
		return errors.New("Missing uuid")
	}
	return nil
}
