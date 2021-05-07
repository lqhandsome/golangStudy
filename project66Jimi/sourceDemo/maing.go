package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println("启动服务错误",err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	b,_ := ioutil.ReadFile("./hello.txt")
	a := fmt.Sprintf(string(b))
	_, _ =fmt.Fprintln(w,string(a))

}