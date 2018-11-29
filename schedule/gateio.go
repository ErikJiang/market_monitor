package schedule

import (

	"bytes"
	"html/template"
	"github.com/JiangInk/market_monitor/extend/api"
	"github.com/JiangInk/market_monitor/extend/email"
	_ "github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

// 行情提醒
func Task1MarketTicker() {
	tick, err := api.GetTicker("EOS")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	// 预警检测
	earlyWarnCheck(tick)
}

// 预警检测
func earlyWarnCheck(tick api.Ticker) {

	// 1. 获取当前任务类型的任务列表

	// 2. 在任务列表中筛选符合预警规则的任务

	// 3. 达到预警条件的任务，批量发送邮件

	type data struct {
		UserName    string
		Email       string
	}

	type rule struct {
		Operator    string  `json:"operator"`   // 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
		WarnPrice   float64 `json:"warnPrice"`  // 预警价格

	}

	isOK := sendWarnNotify(tick)
	if isOK != true {
		log.Info().Msg("send email notify fail")
		return
	}
	log.Info().Msg("send email notify success")
	return

}

func sendWarnNotify(tick api.Ticker) bool {

	subject := "行情预警通知"
	recvEmail := "jiangink@foxmail.com"
	// 生成邮件模板内容
	content, err := genTemplate(tick)
	if err != nil {
		return false
	}
	// 发送邮件
	err = email.SendEmail(subject, recvEmail, content)
	if err != nil {
		return false
	}
	return true
}


type EmailNotifyData struct {
	Title     string
	UserName  string
	Token     string
	LastPrice string
	TickData  api.Ticker
}
// 生成邮件模板
func genTemplate(tick api.Ticker) (string, error) {
	tmpl, err := template.ParseFiles("templates/email.html")
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	data := EmailNotifyData {
		Title:     "行情预警提醒！",
		UserName:  "Jason Marz",
		Token:     "EOS",
		LastPrice: "5.08",
		TickData:  tick,
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, data); err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	return buff.String(), nil
}
