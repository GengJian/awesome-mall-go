package errorCode

var (
	Err10001 = NewErrCode(10001, "查询用户信息失败")
	Err10002 = NewErrCode(10002, "用户名或密码错误")
	Err10003 = NewErrCode(10003, "新增用户失败")
	Err10004 = NewErrCode(10004, "系统异常")
)
