package router

import (
	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/controller/v1"
	"github.com/JiangInk/market_monitor/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.ServerConf.RunMode)
	apiV1 := r.Group("api/v1")
	userController := new(v1.UserController)
	{
		// apiV1.GET("/users", v1.GetUsers)
		// apiV1.POST("/users", v1.AddUser)
		// apiV1.GET("/users/:userId", v1.GetUser)
		// apiV1.PUT("/users/:userId", v1.EditUser)
		// apiV1.DELETE("/users/:userId", v1.RemoveUser)
		
		apiV1.POST("/auth/signup", userController.Signup)
		apiV1.POST("/auth/signin", userController.Signin)
		apiV1.Use(middleware.JWTAuth())
		{
			apiV1.GET("/user", userController.GetUserInfo)
			apiV1.POST("/auth/signout", userController.Signout)
		}
		
	}
	return r
}
