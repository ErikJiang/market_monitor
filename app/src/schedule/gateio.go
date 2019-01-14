package schedule

import (

	"bytes"
	"github.com/ErikJiang/market_monitor/service"
	"html/template"
	"github.com/ErikJiang/market_monitor/extend/api"
	"github.com/ErikJiang/market_monitor/extend/email"
	"github.com/rs/zerolog/log"
	"strconv"
)

// 行情提醒
func Task1MarketTicker() {

	// 预警检测
	earlyWarnCheck()
}

// 预警检测
func earlyWarnCheck() {

	// 获取当前任务类型的任务列表
	taskService := service.TaskService{
		Type: "TICKER",
	}
	list, err := taskService.QueryByType()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msgf("list: %v", list)

	for _, v := range list {
		flag := false
		// 获取预警条目币种对应的当前市场行情
		tick, err := api.GetTicker(v.Token)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
		// 将最新成交价转换为float64类型
		lastPrice, err := strconv.ParseFloat(tick.Last, 64)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}
		// 进行预警规则比较，运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
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

		if flag {
			// 满足预警规则的任务，发送通知邮件
			isok := sendWarnNotify(tick, v)
			if isok != true {
				log.Info().Msgf("send email notify fail, username: %s", v.UserName)
				return
			}
			log.Info().Msgf("send email notify success, username: %s", v.UserName)
		}
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


