package routers

import (
	"github.com/JiangInk/market_monitor/routers/api/v1"
	"github.com/JiangInk/market_monitor/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)	// debug|release|test
	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/users", v1.GetUsers)
		apiV1.POST("/users", v1.AddUser)
		apiV1.GET("/users/:userId", v1.GetUser)
		apiV1.PUT("/users/:userId", v1.EditUser)
		apiV1.DELETE("/users/:userId", v1.RemoveUser)
	}
	return r
}