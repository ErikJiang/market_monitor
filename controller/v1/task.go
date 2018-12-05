package v1

import (
	"github.com/JiangInk/market_monitor/extend/code"
	"github.com/JiangInk/market_monitor/extend/jwt"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/rs/zerolog/log"
	"strconv"
)

// TaskController 用户控制器
type TaskController struct{}

// @Summary 获取任务列表
// @Description 获取当前用户的任务列表
// @Accept json
// @Produce json
// @Tags task
// @Param Authorization header string true "认证 Token 值"
// @Param page query int false "页码"
// @Param pageSize query int false "每页条数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /task [get]
func (tc *TaskController) List(c *gin.Context) {
	log.Info().Msg("enter task list controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	condition := map[string]interface{}{
		"test": "test",
	}
	taskService := service.TaskService{}
	list, count, err := taskService.QueryByPage(condition, int(page), int(pageSize))
	if err !=nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"count": count,
		"list": list,
	})
	return
}

// @Summary 获取任务明细
// @Description 获取当前用户选中任务的明细信息
// @Accept json
// @Produce json
// @Tags task
// @Param Authorization header string true "认证 Token 值"
// @Param taskId path int true "任务ID"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /task/{taskId} [get]
func (tc *TaskController) Retrieve(c *gin.Context) {
	log.Info().Msg("enter task retrieve controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	taskId := c.Param("taskId")
	u64Id, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 通过ID查询任务详情
	taskService := service.TaskService{TaskID: uint(u64Id)}
	task, err := taskService.QueryByID()
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"data": task,
	})
	return
}

type TaskCreateRequest struct {
	TaskType    string  `json:"taskType" binding:"required,oneof= TICKER OTHER"`    // 任务类型
	Operator    string  `json:"operator" binding:"required,oneof= LT LTE GT GTE"`   // 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
	WarnPrice   float64 `json:"warnPrice" binding:"required"`                       // 预警价格
}

type TaskRuleParam struct {
	Operator string
	WarnPrice float64
}

// @Summary 添加新任务
// @Description 当前用户添加新任务
// @Accept json
// @Produce json
// @Tags task
// @Param Authorization header string true "认证 Token 值"
// @Param body body v1.TaskCreateRequest true "创建任务请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 400 {string} json "{"status":400, "code": 4000001, msg:"请求参数有误"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /task [post]
func (tc *TaskController) Create(c *gin.Context) {
	log.Info().Msg("enter task create controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	reqBody := TaskCreateRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	// 整理数据，userId、type、rule ...
	rule := &TaskRuleParam{
	reqBody.Operator,
	reqBody.WarnPrice,
	}
	ruleJson, err := json.Marshal(rule)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	// 创建任务
	taskService := service.TaskService{
		UserID: int(claims.ID),
		Type: reqBody.TaskType,
		Rules: string(ruleJson),
	}
	taskID, err := taskService.StoreTask()
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	log.Debug().Msgf("create task success, taskId: %d", taskID)
	utils.ResponseFormat(c, code.Success, map[string]interface{}{
		"taskId": taskID,
	})

}

// @Summary 修改任务
// @Description 当前用户修改任务明细
// @Accept json
// @Produce json
// @Tags task
// @Param Authorization header string true "认证 Token 值"
// @Param taskId path int true "任务ID"
// @Param body body v1.TaskCreateRequest true "修改任务请求参数"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 400 {string} json "{"status":400, "code": 4000001, msg:"请求参数有误"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /task/{taskId} [put]
func (tc *TaskController) Update(c *gin.Context) {
	log.Info().Msg("enter task update controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	taskId := c.Param("taskId")
	u64Id, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	// 获取请求参数
	reqBody := TaskCreateRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	rule := &TaskRuleParam {
		reqBody.Operator,
		reqBody.WarnPrice,
	}
	ruleJson, err := json.Marshal(rule)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	taskService := service.TaskService{
		UserID: int(claims.ID),
		Type: reqBody.TaskType,
		Rules: string(ruleJson),
	}

	task, msgCode := taskService.UpdateInfo(uint(u64Id))
	if msgCode != nil {
		utils.ResponseFormat(c, msgCode, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, map[string]interface{}{"data": task})
	return

}

// @Summary 删除任务
// @Description 当前用户删除任务
// @Accept json
// @Produce json
// @Tags task
// @Param Authorization header string true "认证 Token 值"
// @Param taskId path int true "任务ID"
// @Success 200 {string} json "{"status":200, "code": 2000001, msg:"请求处理成功"}"
// @Failure 500 {string} json "{"status":500, "code": 5000001, msg:"服务器内部错误"}"
// @Router /task/{taskId} [delete]
func (tc *TaskController) Destroy(c *gin.Context) {
	log.Info().Msg("enter task destroy controller")
	// 获取token信息
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	if claims == nil {
		utils.ResponseFormat(c, code.TokenInvalid, nil)
		return
	}
	// 获取请求参数
	taskId := c.Param("taskId")
	u64Id, err := strconv.ParseUint(taskId, 10, 64)
	if err != nil {
		log.Error().Msg(err.Error())
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	taskService := service.TaskService{}
	err = taskService.RemoveTask(uint(u64Id))
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, map[string]interface{}{"taskId": uint(u64Id)})
	return
}