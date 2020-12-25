package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据
		// username := c.PostForm("username")
		// password := c.PostForm("password")
		// username := c.DefaultPostForm("username","samebody")
		// password := c.DefaultPostForm("xxx","***")

		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}
		password, ok := c.GetPostForm("password")
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	r.Run(":9000")
}
