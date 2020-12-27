package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 访问/index的GET请求会走这一条处理逻辑
	// 路由
	r.HEAD("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "HEAD",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "noroute",
		})
	})

	// 书籍的首页和详情页
	// r.GET("/video/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/video/index",
	// 	})
	// })

	// r.GET("/shop/index", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "/shop/index",
	// 	})
	// })

	// 路由组 
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/00", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/00"})
		})
	}

	r.Group("/shop")
	// 商场首页和详情页

	r.Run(":9000")
}
