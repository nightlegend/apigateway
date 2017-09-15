package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"kubernetes/Godeps/_workspace/src/github.com/ghodss/yaml"
	"log"
	"os"
)

type Configure struct {
	Mysqldbhost     string `json:"mysqldbhost"`
	Mysqldbport     string `json:"mysqldbport"`
	Mysqldbname     string `json:"mysqldbname"`
	Mysqldbusername string `json:"mysqldbusername"`
	Mysqldbpassword string `json:"mysqldbpassword"`
}

func Connect() *sql.DB {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configureFileData, err := ioutil.ReadFile(dir + "/conf/app.conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	var configure Configure
	err = yaml.Unmarshal([]byte(configureFileData), &configure)
	if err != nil {
		log.Fatal(err)
	}

	dsn := configure.Mysqldbusername + ":" + configure.Mysqldbpassword + "@" + "tcp(" + configure.Mysqldbhost +
		":" + configure.Mysqldbport + ")" + "/" + configure.Mysqldbname + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
