package public

import (
	"github.com/nightlegend/apigateway/core/api/docker"
	"github.com/nightlegend/apigateway/core/api/users"
	"github.com/nightlegend/apigateway/core/module"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
)

type LoginInfo struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func PublicApiRouter(router gin.Engine) {
	log.Println("start init public router.......")
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
		if flag {
			c.JSON(http.StatusOK, gin.H{"code": 200, "Message": "Login Successful"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 201, "Message": "Not found your account"})
		}
	})

	log.Println("complete init public router.......")
}
