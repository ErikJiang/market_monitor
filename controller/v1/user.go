package v1

import (
	"fmt"
	"net/http"

	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct{}

var userService = new(service.UserService)

// SignupReqBody 注册请求参数
type SignupReqBody struct {
	Email       string `json:"email" binding:"required"`
	AccountPass string `json:"accountPass" binding:"required"`
	ConfirmPass string `json:"confirmPass" binding:"required"`
}

// Signup 账号注册
func (sc UserController) Signup(c *gin.Context) {
	var reqBody SignupReqBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("> Email: %s\n", reqBody.Email)
	fmt.Printf("> ConfirmPass: %s\n", reqBody.ConfirmPass)

	userID, err := userService.StoreUser(reqBody.Email, reqBody.ConfirmPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": userID})
	return
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
