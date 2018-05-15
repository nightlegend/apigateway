package db

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
)

// Config :
// define mongo struct
type Config struct {
	Mongohost string `json:"mongohost"`
	Mongoport string `json:"mongoport"`
}

// Connectmon :
// connect mongo db.
func Connectmon() *mgo.Session {
	execDirAbsPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/app.conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial("mongodb://" + config.Mongohost + ":" + config.Mongoport)
	if err != nil {
		panic(err)

	}
	return session
}
