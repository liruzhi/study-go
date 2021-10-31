package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"new-gin-blog/business/router"
)

func main() {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello go server"})
	})

	router.Router(e)
	err := e.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
