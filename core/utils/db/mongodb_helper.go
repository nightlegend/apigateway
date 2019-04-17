package db

import (
	"github.com/nightlegend/apigateway/core/utils/consts"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var (
	mongoDB = MongoDB{}
)

// MongoHelper is the API client that performs all operations
type MongoHelper struct {
	Session    *mgo.Session
	DB         string
	Collection string
}

// NewMongoHelper initializes a new mongo helper object.
func NewMongoHelper() MongoHelper {
	return MongoHelper{}
}

// Insert insert one record in db
func (mongoHelper MongoHelper) Insert(ob interface{}) error {
	c := mongoHelper.Session.DB(mongoHelper.DB).C(mongoHelper.Collection)
	// Optional. Switch the session to a monotonic behavior.
	mongoHelper.Session.SetMode(mgo.Monotonic, true)
	err := c.Insert(ob)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// QueryOne query info by condition return one record.
func (mongoHelper MongoHelper) QueryOne(queryCondition interface{}, ob interface{}) (code int, result interface{}) {
	c := mongoHelper.Session.DB(mongoHelper.DB).C(mongoHelper.Collection)
	mongoHelper.Session.SetMode(mgo.Monotonic, true)
	err := c.Find(queryCondition).One(&ob)
	if err != nil {
		return consts.NOACCOUNT, nil
	}
	return consts.SUCCESS, ob
}

// Update is update the entry by latest changed.
// args:
// cName to set the collection name.
// colQuerier to set the query collection condition.
// update to set the update object.
func (mongoHelper MongoHelper) Update(colQuerier interface{}, update interface{}) bool {
	mongoHelper.Session.SetMode(mgo.Monotonic, true)
	c := mongoHelper.Session.DB(mongoHelper.DB).C(mongoHelper.Collection)
	if err := c.Update(colQuerier, update); err != nil {
		log.Println(err)
		return false
	}
	return true
}

// QueryAll is query all record by condition, return a interface.
func (mongoHelper MongoHelper) QueryAll(queryCondition interface{}, ob []interface{}) (result []interface{}) {
	mongoHelper.Session.SetMode(mgo.Monotonic, true)
	c := mongoHelper.Session.DB(mongoHelper.DB).C(mongoHelper.Collection)
	err := c.Find(queryCondition).All(&ob)
	if err != nil {
		log.Println(err)
	}
	return ob
}
