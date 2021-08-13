package main

import (
	"log"

	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// 0.连接mysql
	err := mysql.InitMysql("root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
		return
	}

	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行函数
	route.RegisterRoute(r)
	// 3.监听端口,默认在8080
	err = r.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
