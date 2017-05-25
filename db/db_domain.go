package db

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/swatlabs/taurus/model"
)

//QueryDomain It returns the domain with the uid
func (bc *BoltClient) QueryDomain(domainUID string) (model.Domain, error) {
	domain := model.Domain{}
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DomainBucket"))
		domainBytes := b.Get([]byte(domainUID))
		if domainBytes == nil {
			return fmt.Errorf("No domain found for " + domainUID)
		}
		json.Unmarshal(domainBytes, &domain)
		return nil
	})
	if err != nil {
		return model.Domain{}, err
	}
	return domain, nil
}

//CreateDomain insert objects into the DomainBucket bucket.
func (bc *BoltClient) CreateDomain(domain model.Domain) {
	jsonBytes, _ := json.Marshal(domain)

	// Write the data to the AccountBucket
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DomainBucket"))
		err := b.Put([]byte(domain.UID), jsonBytes)
		return err
	})
}
