package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Sessions : session store
func Sessions(router *gin.Engine) {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}
