package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserInfo 用户信息
type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// c.JSON(http.StatusOK, gin.H{
		// 	"username": username,
		// 	"password": password,
		// })

		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":   "ok",
				"username": u.Username,
				"password": u.Password,
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	})

	r.Run(":9000")
}
