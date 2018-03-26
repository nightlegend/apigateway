package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/router/private"
	"github.com/nightlegend/apigateway/core/router/public"
	"github.com/nightlegend/apigateway/middleware"
)

// Start from here.
func Start() {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	public.APIRouter(router)

	private.APIRouter(router)
	router.Run(":8012")
}
