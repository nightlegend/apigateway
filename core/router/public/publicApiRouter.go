package public

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/api/docker"
	"github.com/nightlegend/apigateway/core/api/users"
	"github.com/nightlegend/apigateway/core/utils/consts"
)

var (
	// User services
	uis  users.UserInfoService
	flag int
)

// APIRouter is route public router
func APIRouter(router *gin.Engine) {
	log.Info("start init public router.......")
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	router.POST("/login", func(c *gin.Context) {
		c.BindJSON(&uis)
		session := sessions.Default(c)
		v := session.Get(uis.USERNAME)
		if v == nil {
			flag = uis.Login()
			session.Set(uis.USERNAME, uis.USERNAME)
			session.Save()
			log.Println("Try login and save session in session store.")
		} else {
			flag = consts.SUCCESS
			log.Println("Have a session in session store.")
		}

		switch flag {
		case consts.SUCCESS:
			c.JSON(http.StatusOK, gin.H{"code": consts.SUCCESS, "Message": "Login Successful", "tooken": ""})
		case consts.NOACCOUNT:
			c.JSON(http.StatusOK, gin.H{"code": consts.NOACCOUNT, "Message": "Not found your account"})
		case consts.SYSERROR:
			c.JSON(http.StatusOK, gin.H{"code": consts.SYSERROR, "Message": "System error!!!"})
		case consts.WRONGPASSWD:
			c.JSON(http.StatusOK, gin.H{"code": consts.WRONGPASSWD, "Message": "Wrong password..."})
		}
	})

	router.POST("/register", func(c *gin.Context) {
		c.BindJSON(&uis)
		result := uis.Register()
		if result {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Welcome " + uis.USERNAME + ",you have login successful!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "Rigster failed "})
		}
	})

	router.POST("/update", func(c *gin.Context) {
		c.BindJSON(&uis)
		res := uis.UpdateUserInfo()
		if res {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": uis.USERNAME + ",you have update information successfully!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "update information failed, please contract admin help..."})
		}
	})

	router.GET("/queryAll", func(c *gin.Context) {
		users := uis.QueryAllAccountInfo()
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Query done, check console", "userList": users})
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

	log.Info("complete init public router.......")
}
