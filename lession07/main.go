package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法1: 使用map
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "hello world!",
		// 	"age":     18,
		// }

		// data := gin.H{
		// 	"name":    "小王子",
		// 	"message": "hello world!",
		// 	"age":     18,
		// }
		// 方法2: 结构体
		type msg struct {
			Name    string `json:"name"`
			Message string
			Age     int
		}
		data := &msg{
			Name:    "小王子",
			Message: "hello world!",
			Age:     10,
		}

		c.JSON(http.StatusOK, data)
	})
	r.Run(":9000")
}
