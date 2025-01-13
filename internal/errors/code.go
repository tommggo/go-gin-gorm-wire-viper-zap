package errors

// ErrCode 定义错误码
type ErrCode struct {
	code    int
	message string
}

// Code 返回错误码
func (e ErrCode) Code() int {
	return e.code
}

// Message 返回错误信息
func (e ErrCode) Message() string {
	return e.message
}

var (
	// 成功
	Success = ErrCode{0, "success"}

	// 系统错误 (1000-1999)
	ErrSystem = ErrCode{1000, "系统错误"}    // 未知系统错误
	ErrDB     = ErrCode{1001, "数据库错误"}   // 数据库错误
	ErrCache  = ErrCode{1002, "缓存错误"}    // 缓存错误
	ErrRPC    = ErrCode{1003, "RPC调用错误"} // RPC错误

	// 业务错误 (2000-2999)
	ErrBusiness = ErrCode{2000, "业务错误"} // 通用业务错误

	// 参数错误 (3000-3999)
	ErrInvalidParam = ErrCode{3000, "参数错误"} // 通用参数错误
)
