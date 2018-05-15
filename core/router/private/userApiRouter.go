package private

import (
	"container/list"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/api/docker"
)

var secrets = gin.H{
	"golang": gin.H{"email": "golang@golang.com", "phone": "123433"},
	"admin":  gin.H{"email": "lose.start.david.guo@gmail.com", "phone": "13798972142", "status": "successful"},
}

// APIRouter : define a private router.
func APIRouter(router *gin.Engine) {

	log.Info("start init private router.......")
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"golang": "golang",
		"admin":  "Password1",
	}))
	// /admin/secrets endpoint
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
			var containers []string
			i := 0
			// = make([]string, containerList.Len())
			for e := containerList.Front(); e != nil; e = e.Next() {
				res1B, _ := json.Marshal(e.Value)
				temp := string(res1B)
				containerStr.PushBack(temp)
				containers[i] = temp
				i++
			}
			str, _ := json.Marshal(containers)
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
	log.Info("complete init private router.......")
}
