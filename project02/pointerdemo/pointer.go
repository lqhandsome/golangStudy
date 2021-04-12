package main

import	(
	"fmt"
)
var (
	i int;f float64
)
func main()  {
	
	var j int = 5
	var x  = &j
	/*
	var xx *int = &j
	xx是一个指针变量
	xx的类型是*int
	ptr存储的值是j的内存地址(&j)
	*/
	var xx *int = &j
	fmt.Println("j的地址=",&j)
	fmt.Printf("%v",*x)
	fmt.Printf("%v",xx)
}