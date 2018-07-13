package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/router/private"
	"github.com/nightlegend/apigateway/core/router/public"
	"github.com/nightlegend/apigateway/middleware"
)

// Define val
var (
	LisAddr string
)

func init() {
	LisAddr = os.Getenv("ADDRESS")
	if LisAddr == "" {
		LisAddr = "0.0.0.0:8080"
	}
}

// Start a api server.Here define two type router(public and private)
func Start(env string) {
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(middleware.CORSMiddleware()) // import CORS
	router.Use(gin.Logger())
	public.APIRouter(router)  //No Permission Validation
	private.APIRouter(router) //Permission Validation
	router.Run(LisAddr)
}
