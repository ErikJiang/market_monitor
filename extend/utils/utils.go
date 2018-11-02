package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/JiangInk/market_monitor/extend/utils/code"
)

// ResponseFormat 返回数据格式化
func ResponseFormat(c *gin.Context, respStatus *code.Code, data interface{}) {
	if respStatus == nil {
		log.Error().Msg("response status param not found!")
		respStatus = code.RequestParamError
	}
	c.JSON(respStatus.Status, gin.H{
		"code": respStatus.Code,
		"msg":  respStatus.Message,
		"data": data,
	})
	return
}

// Md5 计算字符串的md5值
func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}
