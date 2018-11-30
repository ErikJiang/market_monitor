package v1

import (
	"github.com/JiangInk/market_monitor/service"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// AuthController 用户控制器
type AuthController struct{}

type SignupRequest struct {
	Email       string `json:"email" binding:"required,email"`
	AccountPass string `json:"accountPass" binding:"required"`
	ConfirmPass string `json:"confirmPass" binding:"required"`
}

// @Summary 账号注册
// @Description 通过邮箱密码注册账号
// @Accept json
// @Produce json
// @Tags auth
// @ID auth.signup
// @Param body body v1.SignupRequest true "账号注册请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 400 {string} json "{"status":400, "code": 4000001, msg:"请求参数有误"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /auth/signup [post]
func (ac AuthController) Signup(c *gin.Context) {
	log.Info().Msg("enter signup controller")
	reqBody := SignupRequest{}
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

	userService := service.UserService{
		Email: reqBody.Email,
		Password: reqBody.ConfirmPass,
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

type SigninRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=20"`
}

// @Summary 账号登录
// @Description 通过邮箱密码登录账号
// @Accept json
// @Produce json
// @Tags auth
// @ID auth.signin
// @Param body body v1.SigninRequest true "账号登录请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 400 {string} json "{"status":400, "code": 4000001, msg:"请求参数有误"}"
// @Failure 401 {string} json "{"status":401, "code": 4010001, msg:"账号或密码有误"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /auth/signin [post]
func (ac AuthController) Signin(c *gin.Context) {
	log.Info().Msg("enter Signin controller")
	reqBody := SigninRequest{}
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 登录验证
	userService := service.UserService{
		Email: reqBody.Email,
	}
	user, err := userService.QueryByEmail(reqBody.Email)
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
	authService := service.AuthService{
		User: user,
	}
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

// @Summary 账号注销
// @Description 用户账号注销
// @Accept json
// @Produce json
// @Tags auth
// @ID auth.signout
// @Param Authorization header string true "认证 Token 值"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /auth/signout [post]
func (ac AuthController) Signout(c *gin.Context) {
	log.Info().Msg("enter signout controller")
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	log.Debug().Msgf("claims: %v", claims)
	// 销毁 token
	authService := service.AuthService{}
	isOK, err := authService.DestroyToken(claims.Email)
	if err != nil || isOK == false {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	utils.ResponseFormat(c, code.Success, map[string]interface{}{})
	return
}
