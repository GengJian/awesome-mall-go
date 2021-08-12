package services

import (
	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/model"
)

func Login(req model.LoginRequest) bool {
	mysql.DB.QueryRow(`select  * from user where name = ?`, req.Name)
	return false
}
