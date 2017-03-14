package users

import (
	"container/list"
	"docker-registry/core/utils/db"
	"log"
)

var (
	id          int
	name        string
	password    string
	email       string
	phoneNumber string
	image       string
)

type user struct {
	id          int
	name        string
	password    string
	email       string
	phoneNumber string
	image       string
}

func GetAllUser() list.List {
	dbcon := db.Connect()
	rows, err := dbcon.Query("select * from magic.userinfo")
	if err != nil {
		log.Println(err)
	}
	userList := list.New()
	for rows.Next() {
		err := rows.Scan(&id, &name, &password, &email, &phoneNumber, &image)
		if err != nil {
			log.Fatal(err)
		}
		userInfo := user{id, name, password, email, phoneNumber, image}
		userList.PushBack(userInfo)
		log.Println(userInfo)
	}
	return *userList
}
