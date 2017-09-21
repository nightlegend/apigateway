package router

import (
	"apigateway/core/router/private"
	"apigateway/core/router/public"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
)

/*
 *
 * Support for CORS function.
 *
 */
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
	/*
	 *
	 *   Init router
	 *
	 */
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	public.PublicApiRouter(*router)
	private.PrivateApiRouter(*router)
	router.Run(":8012")
}
