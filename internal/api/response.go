package api

import (
	"go-gin-gorm-wire-viper-zap/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`             // 业务码
	Message string      `json:"message"`          // 提示信息
	Detail  string      `json:"detail,omitempty"` // 错误详情（仅开发环境显示）
	Data    interface{} `json:"data,omitempty"`   // 数据
}

// Message 简单消息响应
type Message struct {
	Message string `json:"message"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:    errors.Success.Code(),
		Message: errors.Success.Message(),
		Data:    data,
	})
}

// SuccessMessage 成功消息响应
func SuccessMessage(c *gin.Context, message string) {
	Success(c, &Message{Message: message})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	// 转换为业务错误
	var appErr errors.Error
	if e, ok := err.(errors.Error); ok {
		appErr = e
	} else {
		appErr = errors.Wrap(errors.ErrSystem, err)
	}

	// 确定 HTTP 状态码
	httpStatus := httpStatusFromCode(appErr.Code())

	// 构建响应
	resp := &Response{
		Code:    appErr.Code(),
		Message: appErr.Message(),
	}

	// 在开发环境显示错误详情
	if gin.Mode() == gin.DebugMode {
		resp.Detail = appErr.Detail()
	}

	c.JSON(httpStatus, resp)
}

// httpStatusFromCode 根据业务码确定 HTTP 状态码
func httpStatusFromCode(code int) int {
	switch {
	case code == errors.Success.Code():
		return http.StatusOK
	case code >= 3000 && code < 4000:
		// 参数错误 (3000-3999)
		return http.StatusBadRequest
	case code >= 2000 && code < 3000:
		// 业务错误 (2000-2999)
		return http.StatusBadRequest
	case code >= 1000 && code < 2000:
		// 系统错误 (1000-1999)
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
