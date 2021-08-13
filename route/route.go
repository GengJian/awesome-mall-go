package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*注册路由*/
func RegisterRoute(g *gin.Engine) {
<<<<<<< HEAD
	// 1.第一个GET请求：HelloWorld
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World! Let`s Go!")
	})
=======
	g.POST("/user/login", handler.Login)
	g.POST("/user/register", handler.Register)
>>>>>>> main
}
