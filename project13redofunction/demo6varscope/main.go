package main

import (
	"fmt"
)

var (
	age = 50
	name = "jack"
)
// name := "tom" == var name string  name = "tom" 复制语句只能在函数体内使用
func main()  {
	fmt.Println(age)
	fmt.Println(name)
	var arr [5] int
	arr[0] = 1
	test()
	for i := 0;i < 10;i++{
		for j:= i ;j<10;j++{
			fmt.Println(i,"-",j)
		}
	}
}

func test() {
	age := 60
	name := "tom"
	fmt.Println(age,name)
}