package main

import (
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/models"
	"github.com/JiangInk/market_monitor/router"
	// "github.com/JiangInk/market_monitor/schedule"
)

func main() {
	config.Setup()
	models.Setup()
	// schedule.GateioCronMain()

	r := router.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + strconv.Itoa(config.ServerConf.Port))
}
