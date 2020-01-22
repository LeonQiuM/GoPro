package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()        // 返回默认的路由引擎
	r.GET("/hello", sayHello) // path 以及对应的handle函数

	// RestFul风格
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "GET",
			"status":  0,
			"message": "ok",
		})
	}) // 获取
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "POST",
			"status":  0,
			"message": "ok",
		})
	}) // 添加
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "PUT",
			"status":  0,
			"message": "ok",
		})
	}) // 更新
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "DELETE",
			"status":  0,
			"message": "ok",
		})
	}) // 删除

	r.Run(":8000")
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{ // 返回一个json：Content-Type = "application/json"
		"message": "Hello World!",
	})
}
