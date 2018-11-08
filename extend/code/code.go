package code

import "net/http"

// Code 错误输出数据结构
type Code struct {
	Status  int    `json:"status"`  // HTTP 状态
	Code    int    `json:"code"`    // 业务错误码
	Message string `json:"message"` // 业务错误信息
}

var (
	// Success 请求处理成功
	Success = &Code{http.StatusOK, 2000001, "请求处理成功"}
	// RequestParamError 请求参数错误
	RequestParamError = &Code{http.StatusBadRequest, 4000001, "请求参数错误"}
	// AccountPassUnmatch 该账号原密码不匹配
	AccountPassUnmatch = &Code{http.StatusBadRequest, 4000002, "该账号原密码不匹配"}
	// SignupPassUnmatch 注册两次输入密码不匹配
	SignupPassUnmatch = &Code{http.StatusBadRequest, 4000003, "注册两次输入密码不匹配"}
	// SigninInfoError 账户名或密码有误
	SigninInfoError = &Code{http.StatusUnauthorized, 4010001, "账户名或密码有误"}
	// TokenNotFound 请求未携带Token, 无权访问
	TokenNotFound = &Code{http.StatusUnauthorized, 4010002, "请求未携带Token, 无权访问"}
	// TokenInvalid 无效的Token信息
	TokenInvalid = &Code{http.StatusUnauthorized, 4010003, "无效的Token信息"}
	// ServiceInsideError 服务器内部错误
	ServiceInsideError = &Code{http.StatusInternalServerError, 5000001, "服务器内部错误"}
)
