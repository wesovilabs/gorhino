package connector

import (
	"fmt"
	"github.com/wesovilabs/taurus/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//MongoDB structure
type MongoDB struct {
	session    *mgo.Session
	hostname   string
	database   string
	collection string
	username   string
	password   string
}

//NewMongoDB constructor
func NewMongoDB(hostname, database, collection string) *MongoDB {
	m := new(MongoDB)
	m.hostname = hostname
	m.database = database
	m.collection = collection
	return m
}

//Initialize inititialize mongodb conection
func (mdb *MongoDB) Initialize() error {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mdb.hostname},
		Timeout:  10 * time.Second,
		Database: mdb.database,
		Username: mdb.username,
		Password: mdb.password,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	mdb.session = session
	return nil
}

func (mdb *MongoDB) getDocuments() (result []model.Service, err error) {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		fmt.Println("Error while running the query in mongo")
	}
	return result, err
}

func (mdb *MongoDB) getDocument(field, searchItem string) (result []model.Service, err error) {
	c := mdb.session.DB(mdb.database).C(mdb.collection)
	err = c.Find(bson.M{field: searchItem}).All(&result)
	if err != nil {
		fmt.Println("Error while running the query in mongo")
	}
	return result, err
}

//LoadConfiguration load configuration
func (mdb *MongoDB) LoadConfiguration() (services []model.Service, err error) {
	services, err = mdb.getDocuments()
	if err != nil {
		panic(err)
	}
	return services, err
}
