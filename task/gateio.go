package task

import (
	"log"
	"bytes"
	"html/template"
	"github.com/JiangInk/market_monitor/api"
	"github.com/JiangInk/market_monitor/service"
	_ "github.com/robfig/cron"
)

type EmailPageData struct {
	Title     string
	UserName  string
	Token     string
	LastPrice string
	TickData  api.Ticker
}

func GateioCronMain() {
	// c := cron.New()
	// // 20秒一次轮询
	// c.AddFunc("*/20 * * * * *", marketTicker)
	// c.Start()
	// select {}
	MarketTicker()
}

// 行情提醒
func MarketTicker() {
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
	content := genTemplate(tick)
	// 发送邮件
	err := service.SendEmail(subject, recvEmail, content)
	if err != nil {
		log.Fatal(err)
	}
}


// 生成邮件模板
func genTemplate(tick api.Ticker) string {

	tmpl, err := template.ParseFiles("templates/email.html")
	if err != nil {
		log.Fatal(err)
	}
	data := EmailPageData{
		Title:     "行情预警提醒！",
		UserName:  "Jason Marz",
		Token:     "EOS",
		LastPrice: "5.08",
		TickData:  tick,
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, data); err != nil {
		log.Fatal(err)
	}
	return buff.String()
}
