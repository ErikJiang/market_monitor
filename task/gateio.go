package task

import (
	"log"
	"bytes"
	"html/template"
	"github.com/JiangInk/market_monitor/api"
)

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