package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("监听失败")
		return
	}

}

func sayHello(w http.ResponseWriter, r *http.Request) {

	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("模板创建失败")
		return
	}

	//渲染模板
	//user := User{
	//	"lq",
	//	23,
	//	"man",
	//}
	m1 := map[string]interface{} {
		"name": "小王子",
		"gender" : "男",
		"age": 24,
	}
	err = t.Execute(w, m1)

	if err != nil {
		fmt.Println("模板渲染失败",err)
		return
	}

}
