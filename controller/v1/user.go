package v1

import (
	"net/http"

	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
)

var userService = new(service.UserSerivce)

// Signup 账号注册
func Signup(c *gin.Context) {
	email := c.Request.FormValue("email")
	accountPass := c.Request.FormValue("accountPass")
	confirmPass := c.Request.FormValue("confirmPass")

	if len(email) == 0 || len(accountPass) == 0 || len(confirmPass) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "not found"})
	}
	userService.addUser(email, accountPass)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Signin 账号登录
func Signin() {

}

// Signout 账号注销
func Signout() {

}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"user": "user", "value": "value"})

}

// EditUserInfo 编辑用户信息
func EditUserInfo(c *gin.Context) {

}
