package users

import (
	"apigateway/core/utils/db"
	"container/list"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type Person struct {
	Name  string
	Phone string
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

func Mongotesting() {
	session := db.Connectmon()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)

}
