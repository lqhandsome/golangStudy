package main

import (
	"fmt"
	"strings"
)

var (

)

func init()  {
	
}

//闭包
func main()  {
	add := AddUpper(11)
	res1 := add(100)
	res2 := add(200)
	fmt.Println(res1,res2)

	name := makeSuffix(".jpg")
	fmt.Println("文件名处理后",name("a"))
}

//累加器

func AddUpper(n int)func(int) int {
	return func(x int) int {
		n = n+x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(str string) string {
		if strings.HasSuffix(str,suffix) == false {
			return str + suffix
		}
		return str 
	}
}