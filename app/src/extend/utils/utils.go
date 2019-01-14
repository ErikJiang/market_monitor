package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/ErikJiang/market_monitor/extend/code"
	"os"
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

// IsExist 判断文件或路径是否存在
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// IsPerm 检测文件或路径是否有权限
func IsPerm(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

