package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*注册路由*/
func RegisterRoute(g *gin.Engine) {
	// 1.第一个GET请求：HelloWorld
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World! Let`s Go!")
	})
}
