package users

import (
	"apigateway/core/utils/db"
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

type Users struct {
	User []user
}

func GetAllUser() Users {
	var users Users
	dbcon := db.Connect()
	rows, err := dbcon.Query("select * from magic.userinfo")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &password, &email, &phoneNumber, &image)
		if err != nil {
			log.Fatal(err)
		}
		userInfo := user{id, name, password, email, phoneNumber, image}
		users.User = append(users.User, userInfo)
	}

	// log.Println(users)
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
