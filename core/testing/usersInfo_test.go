package users

import (
	"fmt"
	"github.com/nightlegend/apigateway/core/module"
	"testing"
)

func TestRegister(t *testing.T) {
	var userInfo module.UserInfo
	userInfo.EMAIL = "david.guo@test.com"
	userInfo.PASSWORD = "123456"
	userInfo.USERNAME = "Hello"
	flag := Register(userInfo)
	if flag {
		fmt.Println("Successful")
	}
}
