package schedule

import (

	"bytes"
	"github.com/JiangInk/market_monitor/service"
	"html/template"
	"github.com/JiangInk/market_monitor/extend/api"
	"github.com/JiangInk/market_monitor/extend/email"
	"github.com/rs/zerolog/log"
	"strconv"
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
	taskService := service.TaskService{
		Type: "TICKER",
	}
	list, err := taskService.QueryByType()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msgf("list: %v", list)
	// 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
	var pendingList []service.TaskItem
	lastPrice, err := strconv.ParseFloat(tick.Last, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	for _, v := range list {
		flag := false
		switch v.Operator {
		case "LT":
			if lastPrice < v.WarnPrice {
				flag = true
			}
		case "LTE":
			if lastPrice <= v.WarnPrice {
				flag = true
			}
		case "GT":
			if lastPrice > v.WarnPrice {
				flag = true
			}
		case "GTE":
			if lastPrice >= v.WarnPrice {
				flag = true
			}
		default:
			flag = false
		}
		// 如果满足预警规则条件，则添加到待处理切片中
		if flag {
			pendingList = append(pendingList, v)
		}
	}

	// 达到预警条件的任务，批量发送邮件
	for _, item := range pendingList {
		isok := sendWarnNotify(tick, item)
		if isok != true {
			log.Info().Msgf("send email notify fail, username: %s", item.UserName)
			return
		}
		log.Info().Msgf("send email notify success, username: %s", item.UserName)
	}

}

type EmailNotify struct {
	Title     string
	UserName  string
	Token     string
	LastPrice string
	TickData  api.Ticker
}

// sendWarnNotify 发送预警消息
func sendWarnNotify(tick api.Ticker, task service.TaskItem) bool {
	// 获取邮件模板
	tmpl, err := template.ParseFiles("templates/email.html")
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}

	subject := "行情预警提醒！"
	data := EmailNotify {
		Title:     subject,
		UserName:  task.UserName,
		Token:     task.Token,
		LastPrice: tick.Last,
		TickData:  tick,
	}

	// 渲染模板内容
	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, data); err != nil {
		log.Error().Msg(err.Error())
		return false
	}

	// 发送预警邮箱
	err = email.SendEmail(subject, task.Email, buff.String())
	if err != nil {
		return false
	}
	return true
}


