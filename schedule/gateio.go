package schedule

import (

	"bytes"
	"github.com/JiangInk/market_monitor/service"
	"html/template"
	"github.com/JiangInk/market_monitor/extend/api"
	"github.com/JiangInk/market_monitor/extend/email"
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

type data struct {
	UserName    string
	Email       string
}

type rule struct {
	Operator    string  `json:"operator"`   // 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
	WarnPrice   float64 `json:"warnPrice"`  // 预警价格

}
// 预警检测
func earlyWarnCheck(tick api.Ticker) {

	// 1. 获取当前任务类型的任务列表
	taskService := service.TaskService{
		Type: "TICKER",
	}
	list, err := taskService.QueryByType()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msgf("list: %v", list)
	// LT LTE GT GTE // 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
	var flag1 bool = false
	for _, v := range list {
		//flag1 = false
		switch v.Operator {
		case "LT":
			if tick.Last < v.WarnPrice {
				flag1 = true
			}
		case "LTE":
			if tick.Last <= v.WarnPrice {
				flag1 = true
			}
		case "GT":
			if tick.Last > v.WarnPrice {
				flag1 = true
			}
		case "GTE":
			if tick.Last >= v.WarnPrice {
				flag1 = true
			}
		default:
			flag1 = false
		}
	}
	log.Print(flag1)


	// 2. 在任务列表中筛选符合预警规则的任务

	// 3. 达到预警条件的任务，批量发送邮件


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
