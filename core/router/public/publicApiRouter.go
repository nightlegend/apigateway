package public

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/api/docker"
	"github.com/nightlegend/apigateway/core/api/users"
	"github.com/nightlegend/apigateway/core/module"
	"github.com/nightlegend/apigateway/core/utils"
	"github.com/nightlegend/apigateway/core/utils/consts"
)

// LoginInfo : define login entry
type LoginInfo struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

// APIRouter is route public router
func APIRouter(router *gin.Engine) {
	log.Info("start init public router.......")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	router.GET("/getImageTagInfo", func(c *gin.Context) {
		imageName := c.DefaultQuery("imageType", "") + "/" + c.DefaultQuery("imageName", "")
		imageTag := c.DefaultQuery("imageTag", "")
		str := docker.GetImageTagInfo(imageName, imageTag)
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "imageTagList": string(str)})
	})

	router.GET("/getImageAllTags", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		imageName := c.DefaultQuery("imageName", "")
		str := docker.GetAllTagByImageName(imageName)
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "imageTagList": string(str)})

	})

	router.POST("/register", func(c *gin.Context) {
		var registerInfo module.UserInfo
		c.BindJSON(&registerInfo)
		password := string(registerInfo.PASSWORD)
		registerInfo.PASSWORD = string(utils.Crypted(password))
		result := users.Register(registerInfo)
		if result {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Welcome " + registerInfo.USERNAME + ",you have login sucessful!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "Rigster failed "})
		}
	})

	router.POST("/login", func(c *gin.Context) {
		var loginInfo LoginInfo
		c.BindJSON(&loginInfo)
		flag := users.Login(loginInfo.USERNAME, loginInfo.PASSWORD)
		if flag == consts.SUCCESS {
			c.JSON(http.StatusOK, gin.H{"code": consts.SUCCESS, "Message": "Login Successful", "tooken": ""})
		} else if flag == consts.NOACCOUNT {
			c.JSON(http.StatusOK, gin.H{"code": consts.NOACCOUNT, "Message": "Not found your account"})
		} else if flag == consts.SYSERROR {
			c.JSON(http.StatusOK, gin.H{"code": consts.SYSERROR, "Message": "System error!!!"})
		} else if flag == consts.WRONGPASSWD {
			c.JSON(http.StatusOK, gin.H{"code": consts.WRONGPASSWD, "Message": "Wrong password..."})
		}
	})

	log.Info("complete init public router.......")
}
