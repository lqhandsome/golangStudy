package main

import (
	"fmt"
)

func main()  {
	arr := [6] int{1,2,3,4,5,6}
	//数组定义
	var brr [3]float64
	//数组通过字面量定义
	crr := [6]int{11,22,33,44}
	for index,value := range arr{
		fmt.Println(index,"-",value)
	}
	brr[2] = 9
	for index,value := range brr{
		fmt.Println(index,value)
	}
	for index,value := range crr{
		fmt.Println(index,value)
	}
}