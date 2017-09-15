package router

import (
	"apigateway/core/api/docker"
	// "apigateway/core/api/etcd"
	"apigateway/core/api/users"
	// "apigateway/core/cache"
	"container/list"
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"log"
	"net/http"
)

var secrets = gin.H{
	"golang": gin.H{"email": "golang@golang.com", "phone": "123433"},
	"admin":  gin.H{"email": "david.guo@cargosmart.com", "phone": "13798972142", "status": "successful"},
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			log.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		c.Next()
	}
}

func Start() {

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"golang": "golang",
		"admin":  "Password1",
	}))
	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	authorized.GET("/getDockerAllContainer", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			containerList := docker.GetAllContainer()

			containerStr := list.New()
			var interfaceSlice []string = make([]string, containerList.Len())
			i := 0
			for e := containerList.Front(); e != nil; e = e.Next() {
				res1B, _ := json.Marshal(e.Value)
				temp := string(res1B)
				containerStr.PushBack(temp)
				interfaceSlice[i] = temp
				i++
			}
			str, _ := json.Marshal(interfaceSlice)
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "secret": secret, "containerList": string(str)})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	authorized.GET("/getDockerAllImages", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			str := docker.GetRegistryImages()
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "secret": secret, "images": string(str)})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	authorized.GET("/getImageAllTags", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			str := docker.GetAllTagByImageName(c.Request.FormValue("imageName"))
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "secret": secret, "imageTagList": string(str)})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	/**
	 * [GET description]
	 * @param {[type]} "/getImageTagInfo" [description]
	 * @param {[type]} func(c             *gin.Context) {				user : [description]
	 */
	authorized.GET("/getImageTagInfo", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			str := docker.GetImageTagInfo(c.Request.FormValue("imageName"), c.Request.FormValue("imageTag"))
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "secret": secret, "imageTagList": string(str)})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	authorized.GET("/getUserList", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			userList := users.GetAllUser()
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "secret": secret, "userList": userList})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	router.GET("/", func(c *gin.Context) {
		// str := users.Mongotesting()
		str := "test"
		// cache.SaveLoginSession("david")
		// etcd.SaveAction()
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!" + str})
	})

	router.GET("/getImageTagInfo", func(c *gin.Context) {
		log.Println(">>>>>>>" + c.DefaultQuery("imageTag", "test"))
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

	router.POST("/form_post", func(c *gin.Context) {
		log.Println(c.Request.Header)
		log.Println(c.Request.Body)
		htmlData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)

		}
		// print out
		log.Println(string(htmlData)) //<-- here !
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		log.Println(">>>>>>>>>>>" + message)
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// router.GET("/getImageTagInfo", func(c *gin.Context) {
	// 	log.Println("welcome .....")
	// 	str := docker.GetImageTagInfo(c.Request.FormValue("imageName"), c.Request.FormValue("imageTag"))
	// 	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "imageTagList": string(str)})
	// })
	router.Run(":8012")

}
