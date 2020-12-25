package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("name", "小王子")
		address := c.Query("address")
		name, ok := c.GetQuery("query")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"name": name,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":     http.StatusOK,
			"username": username,
			"address":  address,
		})
	})
	r.Run(":9000")
}
