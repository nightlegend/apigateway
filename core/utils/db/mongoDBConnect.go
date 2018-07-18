package db

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
)

// MongoDB define mongo db instance
type MongoDB struct {
	Mongohost string `json:"mongohost"`
	Mongoport string `json:"mongoport"`
	DBName    string `json:"dbname"`
}

// Connectmon connect mongo db.
func (db MongoDB) Connectmon() (*mgo.Session, string) {
	execDirAbsPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/app.conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(data), &db)
	if err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial("mongodb://" + db.Mongohost + ":" + db.Mongoport)
	if err != nil {
		panic(err)
	}
	return session, db.DBName
}
