package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// Configure :
// define mysql struct
type Configure struct {
	Mysqldbhost     string `json:"mysqldbhost"`
	Mysqldbport     string `json:"mysqldbport"`
	Mysqldbname     string `json:"mysqldbname"`
	Mysqldbusername string `json:"mysqldbusername"`
	Mysqldbpassword string `json:"mysqldbpassword"`
}

// Connect :
// mysql db connect function.
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
