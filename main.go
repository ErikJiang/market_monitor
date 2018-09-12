package main

import (
	"strconv"
	"github.com/JiangInk/market_monitor/routers"
	"github.com/JiangInk/market_monitor/task"
	"github.com/JiangInk/market_monitor/setting"
)


func main() {
	setting.Setup()
	task.GateioCronMain()
	
	r := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":"+ strconv.Itoa(setting.ServerSetting.HttpPort))
}
