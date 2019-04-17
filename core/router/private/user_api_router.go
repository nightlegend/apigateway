package private

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"golang": gin.H{"email": "golang@golang.com", "phone": "123433"},
	"admin":  gin.H{"email": "lose.start.david.guo@gmail.com", "phone": "1111111111", "status": "successful"},
}

// APIRouter : define a private router.
func APIRouter(router *gin.Engine) {

	log.Info("start init private router......")
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
	log.Info("complete init private router.......")
}
