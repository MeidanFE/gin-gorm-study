package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模版
	t, err := template.ParseFiles("./hello.tpl")
	if err != nil {
		fmt.Printf("Parse template failed:%v", err)
		return
	}
	name := "小王子"
	//3.渲染模版
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("reender template start failed err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed,err:%v", err)
		return
	}
}
