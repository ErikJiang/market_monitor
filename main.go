package main

import (
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/database"
	"github.com/JiangInk/market_monitor/router"
	// "github.com/JiangInk/market_monitor/schedule"
)

func main() {
	config.Setup()
	database.Setup()
	// schedule.GateioCronMain()

	r := router.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + strconv.Itoa(config.ServerSetting.HttpPort))
}
