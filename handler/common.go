package handler

import (
	"net/http"

	"github.com/cold-farmer/awesome-mall/common"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, common.BasicResponse{
		Status: 0,
		Data:   data,
		Msg:    "success",
	})
}

func Failed(ctx *gin.Context, err error, data ...interface{}) {
	if v, ok := err.(common.BasicResponse); ok {
		ctx.JSON(http.StatusOK, v)
		return
	}
	var msg string
	if err != nil {
		msg = err.Error()
	} else {
		msg = "unknown error"
	}
	ctx.JSON(http.StatusOK, common.NewRespError(-1, msg, data...))
}
