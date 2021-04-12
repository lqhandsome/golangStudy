package main

import "fmt"
//全局变量声明 不能使用 :=方式 ,可以声明不使用,简短声明的变量只能在函数内部使用
//第一种 如果没有赋值则只能在函数体内赋值
var v1,l1,n1 int
//第二种 int可以省略
var v2 ,l2,n3  = 1,2,3
//第三种
var (
	v3 int = 4
	v4 = "5"
)
func main()  {
	fmt.Println("v2,l2,n3",v2,l2,n3)
	fmt.Println("v3,v4",v3,v4)
}