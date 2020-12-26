package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 请求中获取文件
		f, _ := c.FormFile("f1")
		fmt.Println(f.Filename)
		// dst:=fmt.Sprintf("./%s",f.Filename)
		dst := path.Join("./", f.Filename)
		c.SaveUploadedFile(f, dst)
	})

	r.POST("/uploadMange", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := path.Join("./", file.Filename, string(index))
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run(":9000")
}
