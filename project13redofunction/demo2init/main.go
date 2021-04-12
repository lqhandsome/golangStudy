package main

import (
	"fmt"
	"go_code/project13redofunction/utils"
)
var (
	test = 10
	i = utils.I
)
//init函数,通常可以再init函数中完成初始化
func init(){
	fmt.Println("main.init")
}
func main()  {
	
	fmt.Println("main.main")
}