package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"strings"

	"github.com/JiangInk/market_monitor/api"
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

	auth := smtp.PlainAuth(
		"",
		"jiangink@126.com",
		"cncd[year][name]",
		"smtp.126.com",
	)

	fromName := "Ticker Service"
	from := "jiangink@126.com"
	to := []string{"jiangink@foxmail.com"}
	subject := "行情预警通知"

	const (
		HTMLType  = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\t\n"
		PlainType = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\t\n"
	)
	// 生成邮件模板内容
	tptStr := GenTemplate(tick)

	msg := []byte(
		"To: " + strings.Join(to, ",") + "\r\n" +
			"From: " + fromName + "<" + from + ">\r\n" +
			"Subject: " + subject + "\r\n" +
			HTMLType + "\r\n" + tptStr,
	)

	err := smtp.SendMail(
		"smtp.126.com:25",
		auth,
		from,
		to,
		msg,
	)
	log.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sendEmail end.")
}

type EmailPageData struct {
	Title     string
	UserName  string
	Token     string
	LastPrice string
	TickData  api.Ticker
}

// 生成邮件模板
func GenTemplate(tick api.Ticker) string {

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
