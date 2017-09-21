package users

import (
	"apigateway/core/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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

type Users struct {
	User []user
}

func GetAllUser() Users {
	var users Users
	var userEnity user
	dbcon := db.Connect()
	rows, err := dbcon.Query("select * from magic.userinfo")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&userEnity.id, &userEnity.name, &userEnity.password, &userEnity.email, &userEnity.phoneNumber, &userEnity.image)
		if err != nil {
			log.Fatal(err)
		}
		users.User = append(users.User, userEnity)
	}
	return users
}

func Mongotesting() *Person {
	session := db.Connectmon()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	// err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	result := Person{}
	err := c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	log.Println("Phone:", result.Phone)
	// return result.Phone
	return &result
}
