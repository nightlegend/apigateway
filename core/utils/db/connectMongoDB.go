package db

import (
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mongohost string `json:"mongohost"`
	Mongoport string `json:"mongoport"`

}

func Connectmon() *mgo.Session {
	execDirAbsPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(execDirAbsPath+"/conf/app.conf.yml")
	if err != nil {
		fmt.Println(err)
	}
	var config Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println(err)
	}
	session, err := mgo.Dial("mongodb://"+config.Mongohost+":"+config.Mongoport)
	if err != nil {
		panic(err)
	}
	return session
}
