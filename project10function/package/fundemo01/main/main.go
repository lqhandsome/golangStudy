package main

import (
	"go_code/project10function/package/fundemo01/main/utils" //引入到包所在的文件夹
	"fmt"
)

func main()  {
	res := utils.Cal(2,3)
	fmt.Println(res)
	fmt.Println(utils.Demo())
}