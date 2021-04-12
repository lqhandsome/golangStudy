package main

import "fmt"
//局部变量一旦定义必须使用
func main()  {
//单个声明
	//第一种定义方式
	var age int 
	age = 23
	//第二种方式默认类型
	var name = "lq"
	//第三种方式省略 var
	sex := "man"

//批量声明
//第一种方式
var v1,v2,v3 int  = 1,2,3
//第二种
var n1,n2,n3 = 1,"2","3"
//第三种省略 var
	l1,l2,l3 := 4,5,6
	fmt.Println("age=",age)
	fmt.Println("name=",name)
	fmt.Println("sex=",sex)
	fmt.Println("v1,v2,v3 = ",v1,v2,v3)
	fmt.Println("n1,n2,n3 = ",n1,n2,n3)
	fmt.Println("l1,l2,l3 = ",l1,l2,l3)
}