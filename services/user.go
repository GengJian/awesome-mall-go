package services

import (
	"fmt"
	"github.com/cold-farmer/awesome-mall/common"
	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/errorCode"
	"github.com/cold-farmer/awesome-mall/model"
	"strings"
)

func Login(req model.LoginRequest) common.BasicResponse {

	var user model.User
	err := mysql.DB.QueryRow(`select  id,name,password from user_info where name = ?`, req.Name).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		fmt.Printf("query userInfo errorCode, err:%v\n", err)
		return common.BasicResponse{
			Status: errorCode.Err10001.Code,
			Msg:    errorCode.Err10001.Message,
		}
	}
	if &user == nil || !strings.EqualFold(user.Password, req.Passwd) {
		fmt.Printf("用户名或密码错误")
		return common.BasicResponse{
			Status: errorCode.Err10002.Code,
			Msg:    errorCode.Err10002.Message,
		}
	}
	return common.BasicResponse{
		Status: 200,
		Data:   user,
		Msg:    "success",
	}
}
func Register(req model.RegisterRequest) common.BasicResponse {
	sqlStr := "insert into user_info(first_name, age,sex,password) values (?,?,?,?)"
	res, err := mysql.DB.Exec(sqlStr, req.Name, req.Age, req.Sex, req.Passwd)
	if err != nil {
		fmt.Printf("insert userInfo errorCode, err:%v\n", err)
		return common.BasicResponse{
			Status: errorCode.Err10003.Code,
			Msg:    errorCode.Err10003.Message,
		}
	}
	userId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return common.BasicResponse{
			Status: errorCode.Err10004.Code,
			Msg:    errorCode.Err10004.Message,
		}
	}
	fmt.Printf("insert success, the id is %d.\n", userId)
	return common.BasicResponse{
		Status: 200,
		Data:   userId,
		Msg:    "success",
	}
}
