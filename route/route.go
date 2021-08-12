package route

import (
	"github.com/cold-farmer/awesome-mall/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(g *gin.Engine) {
	g.POST("/user/login", handler.Login)
}
