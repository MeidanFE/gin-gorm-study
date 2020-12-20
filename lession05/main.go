package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./index.tpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "这是index页面"
	// 渲染模版
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模版

	// 解析模版
	t, err := template.ParseFiles("./home.tpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v\n", err)
		return
	}
	msg := "这是home页面"
	// 渲染模版
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./templates/base.tpl", "./templates/index.tpl")
	if err != nil {
		fmt.Printf("parse template failed ,err:%v", err)
		return
	}
	name := "小王子"
	t.ExecuteTemplate(w, "index.tpl", name)
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./templates/base.tpl", "./templates/home.tpl")
	if err != nil {
		fmt.Printf("parse template failed ,err:%v", err)
		return
	}
	name := "七米"
	t.ExecuteTemplate(w, "index.tpl", name)
}

func main() {
	http.HandleFunc("/index", index2)
	http.HandleFunc("/home", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("start server faild,err:%v", err)
		return
	}
}
