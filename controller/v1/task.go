package v1

import "github.com/gin-gonic/gin"

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

}

type TaskCreateRequest struct {
	TaskType    string  `json:"taskType" binding:"required,oneof= TICKER OTHER"`    // 任务类型
	Operator    string  `json:"operator" binding:"required,oneof= LT LTE GT GTE"`   // 运算符 LT:"<" LTE:"<=" GT:">" GTE:">="
	WarnPrice   float64 `json:"warnPrice" binding:"required"`                       // 预警价格
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

}