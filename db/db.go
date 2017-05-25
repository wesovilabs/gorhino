package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/swatlabs/taurus/model"
	"log"
)

//IBoltClient interface
type IBoltClient interface {
	OpenBoltDb()
	QueryDomain(domainUID string) (model.Domain, error)
	CreateDomain(domain model.Domain)
	InitializeBucket()
}

//BoltClient  Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

//OpenBoltDb OpenBoltDb
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("test/taurus.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//InitializeBucket  Creates an "DomainBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) InitializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("DomainBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}
