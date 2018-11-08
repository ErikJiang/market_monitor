package main

import (
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/models"
	"github.com/JiangInk/market_monitor/router"
	"github.com/JiangInk/market_monitor/extend/redis"
	"github.com/JiangInk/market_monitor/extend/logger"
	// "github.com/JiangInk/market_monitor/schedule"
)

func main() {
	// 基本配置初始化
	config.Setup()
	// 日志初始化
	logger.Setup()
	// 数据库初始化
	models.Setup()
	// 缓存初始化
	redis.Setup()
	
	// schedule.GateioCronMain()

	r := router.InitRouter()
	// 服务监听
	r.Run(":" + strconv.Itoa(config.ServerConf.Port))
}