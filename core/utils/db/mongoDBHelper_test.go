package db

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"testing"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var (
	jsonDBFile = flag.String("json_db_file", "../../../testdata/user.json", "A json file containing a list of users")
	Server     dbtest.DBServer
	tempDir, _ = ioutil.TempDir("../../../testdata", "testing")
	users      []UserInfo
)

type UserInfo struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	USERNAME string        `json:"username" binding:"required"`
	PASSWORD string        `json:"password" binding:"required"`
	ROLE     string        `json:"role" binding:"required"`
}

func TestInsert(t *testing.T) {
	loadUsers(*jsonDBFile)
	t.Log(users)
	Server.SetPath(tempDir)
	session := Server.Session()
	defer session.Close()
	// mogoHelper := NewMongoHelper()
	m := &MongoHelper{
		Session:    session,
		DB:         "apigateway",
		Collection: "userInfo",
	}
	for _, v := range users {
		err := m.Insert(bson.M{"username": v.USERNAME, "password": v.PASSWORD})
		if err != nil {
			t.Log(err)
		}
	}
}

func TestQueryOne(t *testing.T) {
	var userInfo UserInfo
	Server.SetPath(tempDir)
	session := Server.Session()
	defer session.Close()
	// mogoHelper := NewMongoHelper()
	mh := &MongoHelper{
		Session:    session,
		DB:         "consul",
		Collection: "acounts",
	}
	// insert one record first
	mh.Insert(bson.M{"username": "admin", "password": "admin"})
	// query test case
	res := mh.QueryOne(bson.M{"username": "admin", "password": "admin"}, userInfo)
	t.Log(res)
}

func TestUpdate(t *testing.T) {
	var userInfo UserInfo
	Server.SetPath(tempDir)
	session := Server.Session()
	defer session.Close()
	// mogoHelper := NewMongoHelper()
	mh := &MongoHelper{
		Session:    session,
		DB:         "consul",
		Collection: "acounts",
	}
	mh.Insert(bson.M{"username": "admin", "password": "admin"})
	res := mh.Update(bson.M{"username": "admin"}, bson.M{"$set": bson.M{"password": "admin1"}})
	if res {
		res := mh.QueryOne(bson.M{"username": "admin"}, userInfo)
		mapstructure.Decode(res, &userInfo)
		t.Log(userInfo.PASSWORD)
		if userInfo.PASSWORD != "admin1" {
			t.Fail()
		}
	} else {
		t.Fail()
	}

}

func loadUsers(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load defaule user list: %v", err)
	}
	if err := json.Unmarshal(file, &users); err != nil {
		log.Fatalf("Failed to load default user list: %v", err)
	}
}
