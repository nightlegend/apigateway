package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nightlegend/apigateway/core/utils/consts"
	"gopkg.in/mgo.v2"
)

var (
	mongoDB = MongoDB{}
)

// MongoHelper is the API client that performs all operations
type MongoHelper struct {
}

// NewMongoHelper initializes a new mongo helper object.
func NewMongoHelper() MongoHelper {
	return MongoHelper{}
}

// Insert insert one record in db
func (mongoHelper MongoHelper) Insert(cName string, ob interface{}) bool {
	session, dbName := mongoDB.Connectmon()
	defer session.Close()
	c := session.DB(dbName).C(cName)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	err := c.Insert(ob)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// QueryOne query info by condition return one record.
func (mongoHelper MongoHelper) QueryOne(cName string, queryCondition interface{}, ob interface{}) (result interface{}) {
	session, dbName := mongoDB.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(cName)

	err := c.Find(queryCondition).One(&ob)
	if err != nil {
		log.Println(err)
		return consts.NOACCOUNT
	}
	return ob
}

// Update is update the entry by latest changed.
// args:
// cName to set the collection name.
// colQuerier to set the query collection condition.
// update to set the update object.
func (mongoHelper MongoHelper) Update(cName string, colQuerier interface{}, update interface{}) bool {
	session, dbName := mongoDB.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(cName)
	if err := c.Update(colQuerier, update); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// QueryAll is query all record by condition, return a interface.
func (mongoHelper MongoHelper) QueryAll(cName string, queryCondition interface{}, ob []interface{}) (result []interface{}) {
	session, dbName := mongoDB.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(cName)
	err := c.Find(queryCondition).All(&ob)
	if err != nil {
		log.Println(err)
	}
	return ob
}
