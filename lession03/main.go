package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./hello.tpl")
	if err != nil {
		fmt.Printf("parse template failed:%v", err)
		return
	}
	u1 := User{
		Name:   "   小王子   ",
		Gender: "男",
		Age:    18,
	}

	// t.Execute(w, u1)
	// 渲染模版

	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}

	hobbyList := []string{
		// "篮球",
		// "足球",
		// "双色球",
	}

	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}

}
