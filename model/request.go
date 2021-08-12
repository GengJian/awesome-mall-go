package model

type LoginRequest struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}
