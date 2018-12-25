package main

import (
	"github.com/JiangInk/market_monitor/extend/validator"
	"github.com/JiangInk/market_monitor/schedule"
	"net/http"
	"context"
	"fmt"
    "os"
    "os/signal"
    "time"
	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/extend/logger"
	"github.com/JiangInk/market_monitor/extend/redis"
	"github.com/JiangInk/market_monitor/models"
	"github.com/JiangInk/market_monitor/router"
	"github.com/rs/zerolog/log"
)

// @title Market Monitor API
// @version 1.0
// @description Market Monitor 简易API文档
// @termsOfService http://swagger.io/terms/

// @contact.name ink
// @contact.url http://jiangink.github.com
// @contact.email jiangink.cn@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8000
// @BasePath /api/v1
func main() {
	// 基本配置初始化
	config.Setup()
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

	serv := &http.Server{
		Addr:			fmt.Sprintf(":%d", config.ServerConf.Port),
		Handler:		router,
		ReadTimeout:	config.ServerConf.ReadTimeout,
		WriteTimeout:	config.ServerConf.WriteTimeout,
		MaxHeaderBytes:	1 << 20,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			log.Error().Msgf("Listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server Shutdown: %v", err)
	}
	log.Info().Msg("Server exiting")
}