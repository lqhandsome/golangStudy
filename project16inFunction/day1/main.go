package main

import (
	"fmt"
)

func main()  {
	num1 := 100
	fmt.Println(num1)

	//new 给值类型分配内存比如 int float64 struct 返回的是指针
	num2 := new(int)
	*num2 = 10
	fmt.Printf("num2=%v num2的&num=%v num2的类型是%T,num2存储的值对应地址的值是%v",num2,&num2,num2,*num2)
}