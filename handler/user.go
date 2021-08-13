package handler

import (
	"net/http"

	"github.com/cold-farmer/awesome-mall/common/log"
	"go.uber.org/zap"

	"github.com/cold-farmer/awesome-mall/model"
	"github.com/cold-farmer/awesome-mall/services"

	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(ctx *gin.Context) {
	var req model.LoginRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Debug(ctx, "xxx", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{"msg": "bad param", "errorCode": err.Error()})
		return
	}
	result, err := services.Login(ctx, req)
	if err != nil {
		Failed(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// Register 注册
func Register(ctx *gin.Context) {
	var req model.RegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "bad param", "errorCode": err.Error()})
		return
	}
	result, err := services.Register(ctx, req)
	if err != nil {
		Failed(ctx, err)
		return
	}
	Success(ctx, result)
}
