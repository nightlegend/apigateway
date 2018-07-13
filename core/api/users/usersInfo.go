package users

import (
	log "github.com/Sirupsen/logrus"
	"github.com/nightlegend/apigateway/core/utils"
	"github.com/nightlegend/apigateway/core/utils/consts"
	"github.com/nightlegend/apigateway/core/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserInfoService export all service about user info action.
type UserInfoService struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
}

// Register handle register action
func (uis UserInfoService) Register() bool {
	session := db.Connectmon()
	defer session.Close()
	c := session.DB("test").C("userInfo")
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	err := c.Insert(uis)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// Login :
// param [userName, password]
// return bool type.
func (uis UserInfoService) Login() int {

	session := db.Connectmon()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("userInfo")

	var userInfo UserInfoService
	err := c.Find(bson.M{"username": uis.USERNAME}).One(&userInfo)
	if err != nil {
		log.Println(err)
		return consts.NOACCOUNT
	}
	if uis.PASSWORD == utils.DeCryptedStr([]byte(userInfo.PASSWORD)) {
		return consts.SUCCESS
	} else if uis.PASSWORD != utils.DeCryptedStr([]byte(userInfo.PASSWORD)) {
		return consts.WRONGPASSWD
	}

	return consts.SYSERROR
}
