package model

import (
	"errors"
)

//DomainSettings structure
type DomainSettings struct {
	DomainType string
	Sandbox    bool
}

//Domain structure
type Domain struct {
	UID      string
	Name     string
	Settings MongoDBConfiguration
}

//Credentials structure
type Credentials struct {
	Username string
	Password string
}

//GitConfiguration structure
type GitConfiguration struct {
	DomainSettings
	credentials Credentials
	uri         string
	branch      string
}

//MongoDBConfiguration structure
type MongoDBConfiguration struct {
	Credentials Credentials
	Hostname    string
	Database    string
	Collection  string
}

//FileSystemConfiguration structure
type FileSystemConfiguration struct {
	DomainSettings
	credentials Credentials
	path        string
}

//IsValid Determines if a domain is valid
func (domain Domain) IsValid() error {
	if domain.UID == "" {
		return errors.New("Missing uuid")
	}
	return nil
}

//ServiceProfile structure
type ServiceProfile struct {
	name  string
	props map[string]interface{}
}

//ServiceVersion structure
type ServiceVersion struct {
	name     string
	profiles []ServiceProfile
}

//Service structure
type Service struct {
	name     string
	versions []ServiceVersion
}
