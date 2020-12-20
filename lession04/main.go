package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, request *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值,要么两个返回值,第二个返回值必须是error类型
	kua := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	// 定义模版

	// 解析模版
	t := template.New("f.tpl") // 创建一个名字是f的模版对象,名字一定要和解析模版对上
	// 告诉模版引擎,我现在多了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	_, err := t.ParseFiles("./f.tpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v", err)
		return
	}

	name := "小王子"
	// 渲染模版
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./t.tpl", "./ul.tpl")
	if err != nil {
		fmt.Printf("parse template failed,v:%v", err)
		return
	}

	name := "小王子"
	// 渲染模版
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
