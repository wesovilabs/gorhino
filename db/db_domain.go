package db

import (
	"github.com/golang/mock/mockgen/model"
	"github.com/boltdb/bolt"
	"fmt"
	"encoding/json"
)

func (bc *BoltClient) QueryDomain(domainUid string) (model.Domain, error) {
	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	domain := model.Domain{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("DomainBucket"))
		// Read the value identified by our domainId supplied as []byte
		domainBytes := b.Get([]byte(domainUid))
		if domainBytes == nil {
			return fmt.Errorf("No domain found for " + domainUid)
		}
		// Unmarshal the returned bytes into the domain struct we created at
		// the top of the function
		json.Unmarshal(domainBytes, &domain)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Domain{}, err
	}
	// Return the Account struct and nil as error.
	return domain, nil
}
