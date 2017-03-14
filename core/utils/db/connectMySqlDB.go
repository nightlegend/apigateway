package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	DB_HOST = "tcp(localhost:3306)"
	DB_NAME = "magic"
	DB_USER = /*"root"*/ "root"
	DB_PASS = /*""*/ "123456"
)

func Connect() *sql.DB {

	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
