package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/cold-farmer/awesome-mall/common"
	"github.com/cold-farmer/awesome-mall/common/log"
	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/model"
)

func Login(ctx context.Context, req model.LoginRequest) (interface{}, error) {
	var user model.User
	err := mysql.DB.QueryRow(`select id,name,password from user_info where name = ? and password = ?`,
		req.Name, req.Passwd).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		log.Debug(ctx, "query userInfo errorCode,", zap.Error(err))
		return nil, common.ErrQueryUser
	}
	if user.Id == 0 {
		log.Debug(ctx, "用户名或密码错误")
		return nil, common.ErrWrongPasswd
	}
	return user, nil
}

func Register(ctx context.Context, req model.RegisterRequest) (interface{}, error) {
	sqlStr := "insert into user_info(first_name, age,sex,password) values (?,?,?,?)"
	res, err := mysql.DB.Exec(sqlStr, req.Name, req.Age, req.Sex, req.Passwd)
	if err != nil {
		log.Debug(ctx, "insert userInfo errorCode", zap.Error(err))
		return nil, common.ErrAddUser
	}
	userId, err := res.LastInsertId()
	if err != nil {
		log.Debug(ctx, "get last insert ID failed", zap.Error(err))
		return nil, common.ErrInternal
	}
	log.Debug(ctx, "insert success", zap.Int64("user_id", userId))
	return userId, nil
}
