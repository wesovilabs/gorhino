package model

import (
	"errors"
	"fmt"
)

type DomainSettings struct {
	domainType string
	public     bool
	sandbox    bool
}

//
type Domain struct {
	uid      string
	name     string
	settings DomainSettings
}

type Credentials struct {
	username string
	password string
}

type GitConfiguration struct {
	credentials Credentials
	uri         string
	branch      string
}

type FileSystemConfiguration struct {
	credentials Credentials
	path        string
}

func (domain Domain) isValid() error {
	fmt.Print("Result:")
	fmt.Print(domain.uid)
	if domain.uid == "" {
		return errors.New("Missing uuid")
	}
	return nil
}
