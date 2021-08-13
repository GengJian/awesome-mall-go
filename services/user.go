package services

import (
	"fmt"
	"strings"

	"github.com/cold-farmer/awesome-mall/common"
	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/model"
)

func Login(req model.LoginRequest) (interface{}, error) {

	var user model.User
	err := mysql.DB.QueryRow(`select  id,name,password from user_info where name = ?`, req.Name).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		fmt.Printf("query userInfo errorCode, err:%v\n", err)
		return nil, common.Err10001
	}
	if &user == nil || !strings.EqualFold(user.Password, req.Passwd) {
		fmt.Printf("用户名或密码错误")
		return nil, common.Err10002
	}
	return user, nil
}

func Register(req model.RegisterRequest) (interface{}, error) {
	sqlStr := "insert into user_info(first_name, age,sex,password) values (?,?,?,?)"
	res, err := mysql.DB.Exec(sqlStr, req.Name, req.Age, req.Sex, req.Passwd)
	if err != nil {
		fmt.Printf("insert userInfo errorCode, err:%v\n", err)
		return nil, common.Err10003
	}
	userId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return nil, common.Err10004
	}
	fmt.Printf("insert success, the id is %d.\n", userId)
	return userId, nil
}
