package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/index" // 请求地址修改了
		r.HandleContext(c)            // 继续后续的处理

	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.Run(":9000")
}
