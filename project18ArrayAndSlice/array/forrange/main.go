package main

import (
	"fmt"
)

func main()  {
	var i, a int = 1,2 

	fmt.Println(i,a)
	s := fmt.Sprintf("%d",i)
	fmt.Println(s)
	// var arr []int 切片
	arr := [3]int{11,22,33}
	updateArr(arr)
	fmt.Println(arr)
	addressArr(&arr)
	fmt.Println(arr)
}

func updateArr(arr [3]int){
	arr[2] = 100
}

func addressArr(arr *[3]int){
	(*arr)[2] =88
}