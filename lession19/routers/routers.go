package routers

import (
	"github.com/gin-gonic/gin"

	"lession19/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	// 静态资源配置
	r.Static("/static", "static")
	// 寻找html模版文件
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.PostTodo)
		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetAllTodo)
		// 查看某一个代办事项
		v1Group.GET("/todo/:id", controller.GetTodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
