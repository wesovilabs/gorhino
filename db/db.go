package db

import (
	"github.com/boltdb/bolt"
	"github.com/wesovilabs/taurus/model"
	"log"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryDomain(domainUid string) (model.Domain, error)
	Seed()
}

// Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}
