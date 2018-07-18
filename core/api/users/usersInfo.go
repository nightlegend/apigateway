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
	mongoHelper = db.MongoHelper{}
)

// UserInfoService export all service about user info action.
type UserInfoService struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
}

// Register handle register action
func (uis UserInfoService) Register() bool {
	return mongoHelper.Insert(uis)
}

// Login param [userName, password], return bool type.
func (uis UserInfoService) Login() int {
	var userInfo UserInfoService
	res := mongoHelper.Query(bson.M{"username": uis.USERNAME}, uis)

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
