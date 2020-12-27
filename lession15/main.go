package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	// panic("121")
	log.Println("index ")
	msg, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})

}

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Next() // 通过
	// c.Abort() // 阻止
	cost := time.Since(start)
	log.Println(cost)
	log.Println("m1 out ...")
}

func m2(c *gin.Context) {
	log.Println("m2 in...")
	c.Set("name", "Simon-bin")
	c.Next()
	log.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 准备工作
	return func(c *gin.Context) {

		// 存放具体逻辑
		// 是否登陆
		if doCheck {
			c.Next()
		} else {
			c.Abort()
		}

	}
}

func main() {
	r := gin.Default()

	r.Use(m1) /// 使用全局中间件 m1

	r.GET("/index", m1, m2, indexHandler)
	r.GET("/shop", m1, func(c *gin.Context) {
		msg, _ := c.Get("name")
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	r.GET("/usr", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	// 路由组权限控制
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "xxGroup",
			})
		})
	}

	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "xx2Group"})
		})
	}
	r.Run(":9000")
}
