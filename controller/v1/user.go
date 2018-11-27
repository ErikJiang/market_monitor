package v1

import (
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// UserController 用户控制器
type UserController struct{}

// @Summary 获取用户信息
// @Description 获取当前登录用户基本信息
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user [get]
func (sc UserController) GetInfo(c *gin.Context) {

	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		utils.ResponseFormat(c, code.Success, map[string]interface{}{ 
			"data": claims,
		})
		return
	}
}

type EditRequest struct {
	Name string `json:"name" binding:"required,max=20"`
}

// @Summary 编辑用户昵称
// @Description 修改当前登录用户昵称
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.EditRequest true "修改用户名称请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user/name [patch]
func (sc UserController) AlterName(c *gin.Context) {
	log.Info().Msg("enter edit info controller")
	// 获取 Token 用户信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := EditRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 修改用户名称
	userService := service.UserService{ UserID: claims.ID }
	updateUser, msgCode := userService.UpdateUserName(reqBody.Name)
	if msgCode != nil {
		utils.ResponseFormat(c, msgCode, nil)
		return
	}
	log.Debug().Msgf("update user result: %v", updateUser)
	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"userId": updateUser.ID,
		"userName": updateUser.UserName,
		"email": updateUser.Email,
	})
}

type PassRequest struct {
	OldPass string `json:"oldPass" binding:"required,max=50"`
	NewPass string `json:"newPass" binding:"required,max=50"`
}

// @Summary 修改用户密码
// @Description 修改当前登录用户密码
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.PassRequest true "修改用户密码请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user/pass [patch]
func (sc UserController) AlterPass(c *gin.Context) {
	log.Info().Msg("enter change pass controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := PassRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 更新用户密码
	userService := service.UserService{ Email:claims.Email }
	updateUser, msgCode := userService.UpdateUserPass(reqBody.OldPass, reqBody.NewPass)
	if msgCode != nil {
		utils.ResponseFormat(c, msgCode, nil)
		return
	}
	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"userId": updateUser.ID,
		"userName": updateUser.UserName,
		"email": updateUser.Email,
	})
}

type AvatarRequest struct {
	Avatar string `json:"avatar"`
}

// @Summary 修改用户头像
// @Description 修改当前登录用户头像
// @Accept multipart/form-data
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param avatar formData file true "用户头像请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user/avatar [patch]
func (sc UserController) AlterAvatar(c *gin.Context) {
	log.Info().Msg("enter change pass controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := AvatarRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 更新用户密码
	// todo ...

	utils.ResponseFormat(c, code.Success, map[string]interface{}{})
}