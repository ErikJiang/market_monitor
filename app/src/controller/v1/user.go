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
func (sc UserController) Retrieve(c *gin.Context) {

	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims != nil {
		utils.ResponseFormat(c, code.Success, map[string]interface{}{ 
			"data": claims,
		})
		return
	}
}

// UserEditRequest 用户编辑请求参数
type UserEditRequest struct {
	Name string `json:"name" binding:"required,max=20"`
}

// @Summary 编辑用户昵称
// @Description 修改当前登录用户昵称
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.UserEditRequest true "修改用户名称请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user/name [patch]
func (sc UserController) AlterName(c *gin.Context) {
	log.Info().Msg("enter user edit info controller")
	// 获取 Token 用户信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := UserEditRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 修改用户名称
	userService := service.UserService{ UserID: claims.ID }
	updateUser, msgCode := userService.UpdateName(reqBody.Name)
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

// UserPassRequest 修改用户密码请求参数
type UserPassRequest struct {
	OldPass string `json:"oldPass" binding:"required,max=50"`
	NewPass string `json:"newPass" binding:"required,max=50"`
}

// @Summary 修改用户密码
// @Description 修改当前登录用户密码
// @Accept json
// @Produce json
// @Tags user
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.UserPassRequest true "修改用户密码请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /user/pass [patch]
func (sc UserController) AlterPass(c *gin.Context) {
	log.Info().Msg("enter user change pass controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := UserPassRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 更新用户密码
	userService := service.UserService{ Email:claims.Email }
	updateUser, msgCode := userService.UpdatePass(reqBody.OldPass, reqBody.NewPass)
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
	log.Info().Msg("enter user change avatar controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取上传文件内容
	file, image, err := c.Request.FormFile("avatar")
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if image == nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 获取头像名称
	uploadService := service.UploadService{}
	avatarName := uploadService.GetImgName(image.Filename)
	fullPath := uploadService.GetImgFullPath()
	// 上传图片格式检测
	if !uploadService.CheckImgExt(avatarName) {
		utils.ResponseFormat(c, code.UploadSuffixError, nil)
		return
	}
	// 上传图片大小检测
	if !uploadService.CheckImgSize(file) {
		utils.ResponseFormat(c, code.UploadSizeLimit, nil)
		return
	}
	err = uploadService.CheckImgPath(fullPath)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	err = c.SaveUploadedFile(image, fullPath+avatarName)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	// 更新用户头像
	userService := service.UserService{ UserID: claims.ID }
	updateUser, msgCode := userService.UpdateAvatar(uploadService.GetImgPath()+avatarName)
	if msgCode != nil {
		utils.ResponseFormat(c, msgCode, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"userId": updateUser.ID,
		"avatar": updateUser.Avatar,
	})
}