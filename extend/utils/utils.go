package utils

import (
	"crypto/sha1"
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

// MakeSha1 计算字符串的 sha1 hash 值
func MakeSha1(source string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(source))
	return hex.EncodeToString(sha1Hash.Sum(nil))
}
