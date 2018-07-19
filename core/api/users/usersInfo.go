package users

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/nightlegend/apigateway/core/utils"
	"github.com/nightlegend/apigateway/core/utils/consts"
	"github.com/nightlegend/apigateway/core/utils/db"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongoDB     = db.MongoDB{}
	mongoHelper = db.NewMongoHelper()
	collection  = "userInfo"
)

// UserInfoService export all service about user info action.
type UserInfoService struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
}

// Register register one new user in db, return a boolean value to make know success or not.
func (uis UserInfoService) Register() bool {
	uis.PASSWORD = string(utils.Crypted(string(uis.PASSWORD))) //encryption password.
	return mongoHelper.Insert(collection, uis)
}

// Login param [userName, password], return a int type, it`s a common code,you can known
// it means by the words.
func (uis UserInfoService) Login() int {
	var userInfo UserInfoService
	res := mongoHelper.Query(collection, bson.M{"username": uis.USERNAME}, uis)

	userInfo, ok := res.(UserInfoService)
	if ok {
		log.Println(userInfo.USERNAME)
	} else {
		mapstructure.Decode(res, &userInfo)
	}

	if uis.PASSWORD == utils.DeCryptedStr([]byte(userInfo.PASSWORD)) {
		return consts.SUCCESS
	} else if uis.PASSWORD != utils.DeCryptedStr([]byte(userInfo.PASSWORD)) {
		return consts.WRONGPASSWD
	}
	return consts.SYSERROR
}

// UpdateUserInfo update user account information
func (uis UserInfoService) UpdateUserInfo() bool {
	uis.PASSWORD = string(utils.Crypted(string(uis.PASSWORD)))
	colQuerier := bson.M{"username": uis.USERNAME}
	update := bson.M{"$set": bson.M{"username": uis.USERNAME, "password": uis.PASSWORD, "email": uis.EMAIL}}
	return mongoHelper.Update(collection, colQuerier, update)
}
