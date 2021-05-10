package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/hello",SayHello)
	http.ListenAndServe(":9000",nil)
}

func SayHello(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"hello world！")
}
