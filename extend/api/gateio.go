package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Ticker struct {
	BaseVolume    string    `json:"baseVolume"`    // 交易量
	Last          string    `json:"last"`          // 最新成交价
	HighestBid    string    `json:"highestBid"`    // 买方最高价
	High24hr      string    `json:"high24hr"`      // 24小时最高价
	LowestAsk     string    `json:"lowestAsk"`     // 卖方最低价
	Low24hr       string    `json:"low24hr"`       // 24小时最低价
	PercentChange string    `json:"percentChange"` // 涨跌百分比
	QuoteVolume   string    `json:"quoteVolume"`   // 兑换货币交易量
}

func GetTicker(token string) (tick Ticker, err error) {
	log.Info().Msgf("enter getTicker, token: %s", token)
	var url = "https://data.gateio.io/api2/1/ticker/"+ strings.ToLower(token) +"_usdt"
	resp, err := http.Get(url)
	if err != nil {
		return tick, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tick, err
	}

	json.Unmarshal(body, &tick)

	// 展示行情信息
	colorPrinter(url, token, tick)

	return tick, nil
}

// 行情信息打印
func colorPrinter(url, token string, tick Ticker) {
	color.Red("【当前市场行情】")
	color.Magenta("数据来源:\t%s", url)
	color.Green("当前币种:\t%s", token)
	color.Blue("当前时间:\t%s", time.Now().Format("2006-01-02 15:04:05"))
	color.Cyan("交易量:\t\t%s", tick.BaseVolume)
	color.Cyan("最新成交价:\t%s", tick.Last)
	color.Cyan("买方最高价:\t%s", tick.HighestBid)
	color.Cyan("24小时最高价:\t%s", tick.High24hr)
	color.Cyan("卖方最低价:\t%s", tick.LowestAsk)
	color.Cyan("24小时最低价:\t%s", tick.Low24hr)
	color.Cyan("涨跌百分比:\t%s", tick.PercentChange)
	color.Cyan("兑换货币交易量:\t%s", tick.QuoteVolume)
	color.Yellow("----------------------------------")
}
