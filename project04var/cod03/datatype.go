package main

import "fmt"

var (
	v3 = "4"
	v4 = "5"
)
//基本数据类型
	/*
	数值型 
		整型 int,int8(8位 范围 -127 ~ 126) uint8(无符号),byte(一个字节)
		浮点型 float32(32位),float64
	字符型(ascii) a,b,c,d 单个字符
	布尔型 true or false
	字符串
	*/
//派生型
	/*
	指针
	数组
	结构体 类似于class
	管道
	函数
	切片
	接口
	map
	*/
func main()  {
	fmt.Println("v3,v4",v3+v4)
}