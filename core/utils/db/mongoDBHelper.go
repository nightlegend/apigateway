package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nightlegend/apigateway/core/utils/consts"
	"gopkg.in/mgo.v2"
)

var (
	mongoDB = MongoDB{}
)

// MongoHelper object
type MongoHelper struct {
}

// NewMongoHelper initializes a new mongo helper.
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

// Query query info by condition
func (mongoHelper MongoHelper) Query(cName string, queryCondition interface{}, ob interface{}) (result interface{}) {
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

// Update update entry
func (mongoHelper MongoHelper) Update(cName string, colQuerier interface{}, update interface{}) bool {
	session, dbName := mongoDB.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(dbName).C(cName)
	if err := c.Update(colQuerier, update); err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}
