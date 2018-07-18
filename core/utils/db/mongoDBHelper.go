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

// Insert insert one record in db
func (mongoHelper MongoHelper) Insert(ob interface{}) bool {
	session := mongoDB.Connectmon()
	defer session.Close()
	c := session.DB("test").C("userInfo")
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
func (mongoHelper MongoHelper) Query(queryCondition interface{}, ob interface{}) (result interface{}) {
	session := mongoDB.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("userInfo")

	err := c.Find(queryCondition).One(&ob)
	if err != nil {
		log.Println(err)
		return consts.NOACCOUNT
	}
	return ob
}
