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

// @Summary 编辑用户信息
// @Description 修改当前登录用户基本信息(如：账户昵称)
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.EditRequest true "修改用户信息请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user [put]
func (sc UserController) EditInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		utils.ResponseFormat(c, code.Success, map[string]interface{}{
			"data": claims,
		})
		return
	}
	log.Debug().Msg("enter edit info controller")
	reqBody := EditRequest{}

	// 获取待修改的用户名称
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	// 检测用户名称是否被使用
	userService := service.UserService{
		Name: reqBody.Name,
	}
	user, err := userService.QueryUserByName(reqBody.Name)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if user != nil {
		utils.ResponseFormat(c, code.AccountNameExist, nil)
		return
	}

	// 更新用户名称
	// todo
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
// @Router /user [patch]
func (sc UserController) AlterPass(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		utils.ResponseFormat(c, code.Success, map[string]interface{}{
			"data": claims,
		})
		return
	}
}
