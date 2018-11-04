package v1

import (
	"net/http"

	"github.com/JiangInk/market_monitor/service"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/extend/utils/code"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// UserController 用户控制器
type UserController struct{}

var userService = new(service.UserService)

// Signup 账号注册
func (sc UserController) Signup(c *gin.Context) {
	log.Info().Msg("enter signup controller")
	reqBody := struct {
		Email       string `json:"email" binding:"required,email"`
		AccountPass string `json:"accountPass" binding:"required"`
		ConfirmPass string `json:"confirmPass" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	log.Debug().Msgf("email param: %s", reqBody.Email)
	log.Debug().Msgf("confirmPass param: %s", reqBody.ConfirmPass)

	if reqBody.AccountPass != reqBody.ConfirmPass {
		utils.ResponseFormat(c, code.SignupPassUnmatch, nil)
		return
	}

	userID, err := userService.StoreUser(reqBody.Email, reqBody.ConfirmPass)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	log.Info().Msgf("signup controller result userId: %d", userID)

	utils.ResponseFormat(c, code.Success, map[string]uint{ "userId": userID })
	return
}

// Signin 账号登录
func (sc UserController) Signin(c *gin.Context) {
	log.Info().Msg("enter Signin controller")
	reqBody := struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	// find user info
	user, err := userService.QueryUser()
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	log.Info().Msgf("find user result: %v", user)

	// user info vs request params todo 

	utils.ResponseFormat(c, code.Success, map[string]interface{}{ "user": user })
	return
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
