package errorx

const (
	defaultErrorCode = 10001
	InvalidParams    = 10002 //参数不正确

	defaultMessage = "服务繁忙"
)

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type ErrorCodeResq struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// ErrorCode 实现 error接口
func (e ErrorCode) Error() string {
	return e.Message
}

func NewErrCode(code int, msg string) error {
	return &ErrorCode{
		Code:    code,
		Message: msg,
	}
}

func NewDefaultErrCode() error {
	return &ErrorCode{
		Code:    defaultErrorCode,
		Message: defaultMessage,
	}
}

func (e ErrorCode) Response() *ErrorCodeResq {
	return &ErrorCodeResq{
		Code:    e.Code,
		Message: e.Message,
	}
}
