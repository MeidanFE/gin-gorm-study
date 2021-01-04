package controller

import (
	"log"
	"net/http"

	"lession19/model"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func PostTodo(c *gin.Context) {
	var todo model.Todo
	// 1.请求中拿到todo 对象
	c.BindJSON(&todo)
	// 2.存入数据库
	if err := model.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		//3.返回响应
		c.JSON(http.StatusOK, todo)
	}

}

func GetAllTodo(c *gin.Context) {
	// 查找todo 表里所有数据
	todoList, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")

	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	log.Println("updateTodo:" + id)
	todo, err := model.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)

	if err = model.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
	}

	if err := model.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
