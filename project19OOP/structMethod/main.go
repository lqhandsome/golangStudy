 package main

import "fmt"

 type A struct {
	 Age  int
	 Name string
 }
 //定义一个方法
 func (a A)modify(){
	 a.Name = "handsome"
 }
func main() {


	a := A	{
		23,
		"lq",
	}
	a.modify()
	//基本介绍
	fmt.Println(a)
}

func a(a int,arges... int)   {
	 fmt.Println(arges[0])
	 fmt.Println(a)
 }