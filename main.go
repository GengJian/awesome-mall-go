package main

import (
	"log"

	"github.com/cold-farmer/awesome-mall/dao/mysql"
	"github.com/cold-farmer/awesome-mall/route"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	route.RegisterRoute(g)
	// 连接mysql
	err := mysql.InitMysql("root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println(err)
		return
	}

	err = g.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
