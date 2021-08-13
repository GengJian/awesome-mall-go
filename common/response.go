package common

import (
	"strconv"
)

type BasicResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

var (
	ErrQueryUser   = NewRespError(10001, "查询用户信息失败")
	ErrWrongPasswd = NewRespError(10002, "用户名或密码错误")
	ErrAddUser     = NewRespError(10003, "新增用户失败")
	ErrInternal    = NewRespError(10004, "系统异常")
)

func NewRespError(status int, msg string, data ...interface{}) error {
	resp := BasicResponse{
		Status: status,
		Data:   nil,
		Msg:    msg,
	}
	if len(data) != 0 {
		resp.Data = data
	}
	return resp
}

func (r BasicResponse) Error() string {
	return strconv.Itoa(r.Status) + ":" + r.Msg
}
