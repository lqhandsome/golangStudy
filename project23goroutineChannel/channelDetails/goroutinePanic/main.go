package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	go sayHello()
	for  {
		time.Sleep(time.Second )
		fmt.Println(1)
	}
}
func test(){
	defer func() {
		if err := recover();err != nil{
			fmt.Println("程序运行错误")
		}
	}()
	var m map[int]string
	m[1] = "lq"
}

func sayHello(){
	for  {
		time.Sleep(time.Second * 2)
		fmt.Println("hello")
	}
}