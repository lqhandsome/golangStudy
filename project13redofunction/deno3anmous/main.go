package main

import (
	"fmt"
)

var (
	//全面匿名函数
	fun = func (n1 int,n2 int)int {
		return n1 * n2
	}(10,20)
)

func init()  {
	
}

func main()  {
	//匿名函数调用方式一,直接使用
	res1 := func(n1 int,n2 int) int {
		return n1 + n2
	}(10,20)
	fmt.Println(res1)

	//匿名函数调用方式二,传给一个变量
	niming := func(n1 int,n2 int) int {
		return n1 - n2
	}
	//全局匿名函数
	fmt.Println(niming(10,20))
	//普通函数的传入变量
	ff := test

	fmt.Println(ff(20))
	fmt.Println(fun)
}

func test (n1 int )int {
	return n1 + 2
}