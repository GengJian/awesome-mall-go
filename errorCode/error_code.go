package errorCode

type ErrCode struct {
	Code    int
	Message string
}

func NewErrCode(code int, msg string) *ErrCode {
	e := &ErrCode{
		Code:    code,
		Message: msg,
	}
	return e
}
