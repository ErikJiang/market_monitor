package main

import (
	"github.com/ErikJiang/market_monitor/extend/validator"
	"github.com/ErikJiang/market_monitor/schedule"
	"fmt"
	"github.com/ErikJiang/market_monitor/extend/conf"
	"github.com/ErikJiang/market_monitor/extend/logger"
	"github.com/ErikJiang/market_monitor/extend/redis"
	"github.com/ErikJiang/market_monitor/models"
	"github.com/ErikJiang/market_monitor/router"
)

// @title Market Monitor API
// @version 1.0
// @description Market Monitor 简易API文档
// @termsOfService http://swagger.io/terms/

// @contact.name Erik
// @contact.url http://ErikJiang.github.com
// @contact.email jiangink.cn@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /api/v1
func main() {
	// 基本配置初始化
	conf.Setup()
	// 日志初始化
	logger.Setup()
	// 数据库初始化
	models.Setup()
	// 缓存初始化
	redis.Setup()
	// 验证器初始化
	validator.Setup()
	// 调度任务初始化
	schedule.Setup()

	router := router.InitRouter()

	router.Run(fmt.Sprintf(":%d", conf.ServerConf.Port))
}