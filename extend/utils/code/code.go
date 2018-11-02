package code

import "net/http"

// CodeData 错误输出数据结构
type CodeData struct {
	Status  int    `json:"status"`  // HTTP 状态
	Code    int    `json:"code"`    // 业务错误码
	Message string `json:"message"` // 业务错误信息
}

var (
	// Success 请求处理成功
	Success = &CodeData{http.StatusOK, 2000001, "请求处理成功"}
	// RequestParamError 请求参数错误
	RequestParamError = &CodeData{http.StatusBadRequest, 4000001, "请求参数错误"}
	// AccountPassUnmatch 该账号原密码不匹配
	AccountPassUnmatch = &CodeData{http.StatusBadRequest, 4000002, "该账号原密码不匹配"}
	// SignupPassUnmatch 注册两次输入密码不匹配
	SignupPassUnmatch = &CodeData{http.StatusBadRequest, 4000003, "注册两次输入密码不匹配"}
	// ServiceInsideError 服务器内部错误
	ServiceInsideError = &CodeData{http.StatusInternalServerError, 5000001, "服务器内部错误"}
)
