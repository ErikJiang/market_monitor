package v1

import (
	"github.com/JiangInk/market_monitor/service"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// UserController 用户控制器
type UserController struct{}

// 用户相关服务
var userService = new(service.UserService)

// 认证相关服务
var authService = new(service.AuthService)

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
		Password string `json:"password" binding:"required,max=20"`
	}{}
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	// 登录验证
	user, err := userService.QueryUserByEmail(reqBody.Email)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if user == nil || user.Password != utils.MakeSha1(reqBody.Email+reqBody.Password) {
		utils.ResponseFormat(c, code.SigninInfoError, nil)
		return
	}
	
	// 生成 Token
	token, err := authService.GenerateToken(*user)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, map[string]interface{}{ 
		"userId": user.ID,
		"userName": user.UserName,
		"email": user.Email,
		"token": token,
	})
	return
}

// Signout 账号注销
func (sc UserController) Signout(c *gin.Context) {
	log.Info().Msg("enter signout controller")
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	log.Debug().Msgf("claims: %v", claims)
	// 销毁 token
	isOK, err := authService.DestroyToken(claims.Email)
	if err != nil || isOK == false {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	utils.ResponseFormat(c, code.Success, map[string]interface{}{})
	return
}

// GetUserInfo 获取用户信息
func (sc UserController) GetUserInfo(c *gin.Context) {

	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		utils.ResponseFormat(c, code.Success, map[string]interface{}{ 
			"data": claims,
		})
		return
	}
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
