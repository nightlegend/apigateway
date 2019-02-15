package public

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	mygrpc "github.com/nightlegend/apigateway/core/grpc"
	pb "github.com/nightlegend/apigateway/core/grpc/services"
	ws "github.com/nightlegend/apigateway/core/socketio"

	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/api/users"
	"github.com/nightlegend/apigateway/core/utils/consts"
)

var (
	uis  users.UserInfoService
	flag int
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	log.Info("start init public router.......")
	/*
		// TO-DO: cache store solution
		store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
		if err != nil {
			return
		}
		router.Use(sessions.Sessions("mysession", store))

	*/

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	router.POST("/api/users/login", func(c *gin.Context) {
		c.BindJSON(&uis)
		/*
			// TO-DO: cache user login session.
			session := sessions.Default(c)
			if session.Get(uis.USERNAME) == nil {
				flag = uis.Login()
				session.Set(uis.USERNAME, uis.USERNAME)
				session.Save()
				log.Println("Try login and save session in session store.")
			} else {
				flag = consts.SUCCESS
				log.Println("Have a session in session store.")
			}
		*/
		flag = uis.Login()
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

	router.POST("/api/users/register", func(c *gin.Context) {
		c.BindJSON(&uis)
		if uis.Register() {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Welcome " + uis.USERNAME + ",you have login successful!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "Rigster failed "})
		}
	})

	router.POST("/api/users/update", func(c *gin.Context) {
		c.BindJSON(&uis)
		if uis.UpdateUserInfo() {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": uis.USERNAME + ",you have update information successfully!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "update information failed, please contract admin help..."})
		}
	})

	router.GET("/api/users/queryAll", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Query done, check console", "userList": uis.QueryAllAccountInfo()})
	})

	// export a api for call grpc methods
	router.GET("/api/grpc/demo", func(c *gin.Context) {
		conn, client, err := mygrpc.APIClient()
		defer conn.Close()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err})
		}
		resp := mygrpc.Testing(client, &pb.Request{Id: 1, Msg: "grpc call testinng"})
		c.JSON(http.StatusOK, gin.H{"resp": resp})
	})

	// web socket expose
	router.GET("/ws", func(c *gin.Context) {
		ws.WShandler(c.Writer, c.Request)
	})

	log.Info("complete init public router.......")

}
