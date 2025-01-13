package errors

import "fmt"

// Error 定义业务错误接口
type Error interface {
	error
	Code() int       // 错误码
	Message() string // 错误信息
	Detail() string  // 详细错误信息
	Cause() error    // 原始错误
}

// AppError 实现 Error 接口
type AppError struct {
	errCode ErrCode // 错误码
	detail  string  // 详细错误信息
	cause   error   // 原始错误
}

// New 创建新的错误
func New(errCode ErrCode) Error {
	return &AppError{
		errCode: errCode,
	}
}

// Wrap 包装已有错误
func Wrap(errCode ErrCode, cause error) Error {
	return &AppError{
		errCode: errCode,
		detail:  cause.Error(),
		cause:   cause,
	}
}

// Code 实现 Error 接口
func (e *AppError) Code() int {
	return e.errCode.Code()
}

// Message 实现 Error 接口
func (e *AppError) Message() string {
	return e.errCode.Message()
}

// Detail 实现 Error 接口
func (e *AppError) Detail() string {
	return e.detail
}

// Cause 实现 Error 接口
func (e *AppError) Cause() error {
	return e.cause
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.detail != "" {
		return fmt.Sprintf("[%d] %s: %s", e.Code(), e.Message(), e.detail)
	}
	return fmt.Sprintf("[%d] %s", e.Code(), e.Message())
}
