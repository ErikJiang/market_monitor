package v1

import (
	"net/http"

	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct{}

var userService = new(service.UserService)

// Signup 账号注册
func (sc UserController) Signup(c *gin.Context) {
	email := c.Request.FormValue("email")
	accountPass := c.Request.FormValue("accountPass")
	confirmPass := c.Request.FormValue("confirmPass")

	if len(email) == 0 || len(accountPass) == 0 || len(confirmPass) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "request param error"})
	}
	userID, err := userService.StoreUser(email, accountPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "user signup fail"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": userID})
}

// Signin 账号登录
func (sc UserController) Signin(c *gin.Context) {

}

// Signout 账号注销
func (sc UserController) Signout(c *gin.Context) {

}

// GetUserInfo 获取用户信息
func (sc UserController) GetUserInfo(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"user": "user", "value": "value"})

}

// EditUserInfo 编辑用户信息
func (sc UserController) EditUserInfo(c *gin.Context) {

}

// 修改用户密码
func (sc UserController) alterPasswd(c *gin.Context) {

}

// 修改用户邮箱
func (sc UserController) alterEmail(c *gin.Context) {

}
