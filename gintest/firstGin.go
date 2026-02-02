package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
			"status":  "success",
		})
	})

	// 启动服务 默认在8080 端口
	r.Run()
}
