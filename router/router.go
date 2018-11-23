package router

import (
	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/controller/v1"
	_ "github.com/JiangInk/market_monitor/docs"
	"github.com/JiangInk/market_monitor/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-contrib/cors"
	"time"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.ServerConf.RunMode)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool { return true },
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("api/v1")
	authController := new(v1.AuthController)
	{
		// 账号注册
		apiV1.POST("/auth/signup", authController.Signup)
		// 账号登录
		apiV1.POST("/auth/signin", authController.Signin)
		userController := new(v1.UserController)
		apiV1.Use(middleware.JWTAuth())
		{
			// 账户注销
			apiV1.POST("/auth/signout", authController.Signout)
			// 查看用户信息
			apiV1.GET("/user", userController.GetInfo)
		}
		
	}
	return r
}
