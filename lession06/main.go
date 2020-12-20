package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 静态文件
// html

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.Static("/static", "./statics")
	// r.LoadHTMLFiles("templates/posts/index.tpl", "templates/users/index.tpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "liwenzhou.com",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index.tpl",
		})
	})

	r.Run(":9090")
}
