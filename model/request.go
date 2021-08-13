package model

type LoginRequest struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}

type RegisterRequest struct {
	Name        string `json:"name"`
	Passwd      string `json:"passwd"`
	Age         int    `json:"age"`
	Sex         int    `json:"sex"`
	MobilePhone string `json:"mobilePhone"`
}
