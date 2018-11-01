package v1

import (
	"net/http"

	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
	log.Info().Msg("enter signup controller")
	var reqBody SignupReqBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request params is error"})
		return
	}

	log.Debug().Msgf("email param: %s", reqBody.Email)
	log.Debug().Msgf("confirmPass param: %s", reqBody.ConfirmPass)

	if reqBody.AccountPass != reqBody.ConfirmPass {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user password unmatch!"})
		return
	}

	userID, err := userService.StoreUser(reqBody.Email, reqBody.ConfirmPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info().Msgf("signup controller result userId: %d", userID)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "userId": userID})
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
