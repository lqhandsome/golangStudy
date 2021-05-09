package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//用原生的模板引擎

func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("监听失败")
		return
	}

}

func sayHello(w http.ResponseWriter,r *http.Request) {

	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("模板创建失败")
		return
	}

	//渲染模板
	name := "小胖子"
	err = t.Execute(w,name)
	if err != nil {
		fmt.Println("模板渲染失败")
		return
	}

}
