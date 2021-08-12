package handler

import (
	"net/http"

	"github.com/cold-farmer/awesome-mall/model"
	"github.com/cold-farmer/awesome-mall/services"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var req model.LoginRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "bad param", "error": err.Error()})
		return
	}
	ok := services.Login(req)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"msg": "can not login"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}
