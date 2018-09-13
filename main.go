package main

import (
	"strconv"

	"github.com/JiangInk/market_monitor/config"
	"github.com/JiangInk/market_monitor/models"
	"github.com/JiangInk/market_monitor/routers"
	// "github.com/JiangInk/market_monitor/task"
)

func main() {
	config.Setup()
	models.Setup()
	// task.GateioCronMain()

	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + strconv.Itoa(config.ServerSetting.HttpPort))
}
