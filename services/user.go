package services

import (
	"fmt"
	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/model"
	"strings"
)

func Login(req model.LoginRequest) bool {

	var user model.User
	err := mysql.DB.QueryRow(`select  id,name,password from user_info where name = ?`, req.Name).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		fmt.Printf("query userInfo error, err:%v\n", err)
		return false
	}
	if &user == nil || !strings.EqualFold(user.Password, req.Passwd) {
		fmt.Printf("用户名或密码错误")
		return false
	}
	return true
}
func Register(req model.RegisterRequest) bool {
	sqlStr := "insert into user_info(name, age,sex,password) values (?,?,?,?)"
	res, err := mysql.DB.Exec(sqlStr, req.Name, req.Age, req.Sex, req.Passwd)
	if err != nil {
		fmt.Printf("insert userInfo error, err:%v\n", err)
		return false
	}
	userId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return false
	}
	fmt.Printf("insert success, the id is %d.\n", userId)
	return true
}
