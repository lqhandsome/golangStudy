package main

import (
	"fmt"
)

func main()  {
	fmt.Println(test(2))
	and , cha := manyParms(5,2)
	fmt.Println(and,cha)
}
//单个参数单个返回值
func test(a int) int {
	return a * 2
}

//多个参数.多个返回值

func manyParms(a int,b int) (int,int) {
	return a + b, a - b
}