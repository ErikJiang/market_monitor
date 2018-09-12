package main

import (
	"log"
	"github.com/JiangInk/market_monitor/api"
	"github.com/JiangInk/market_monitor/service"
	"github.com/JiangInk/market_monitor/task"
	_ "github.com/robfig/cron"
)

func main() {
	// c := cron.New()
	// // 20秒一次轮询
	// c.AddFunc("*/20 * * * * *", marketTicker)
	// c.Start()
	// select {}
	marketTicker()
}

// 行情提醒
func marketTicker() {
	tick, err := api.GetTicker("EOS")
	if err != nil {
		log.Fatal(err)
	}

	// 预警检测
	earlyWarnCheck(tick)
}

// 预警检测
func earlyWarnCheck(tick api.Ticker) {

	// 1. 检测判断逻辑

	// 2. 达到预警条件，发送邮件
	sendEmail(tick)

}

func sendEmail(tick api.Ticker) {
	log.Println("enter sendEmail.")

	subject := "行情预警通知"
	recvEmail := "jiangink@foxmail.com"
	// 生成邮件模板内容
	content := task.GenTemplate(tick)
	// 发送邮件
	err := service.SendEmail(subject, recvEmail, content)
	if err != nil {
		log.Fatal(err)
	}
}

