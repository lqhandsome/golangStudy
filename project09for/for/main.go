package main

import (
	"fmt"
)
func main()  {
	var i int = 0
	for ; i < 1; i++ {
		fmt.Println("你好")
	}
	for {
		fmt.Println("hello world")
		break
	}
	//字符串遍历
	str := "hello world刘强"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c%d\n",str2[i],i)
	}
	//使用range
	str = "s刘强"
	for index,val := range str {
		fmt.Printf("%d-%c,%T\n",index,val,val)
	}
	arr :=[2]string{"liuq","two"}
	for index,val := range arr {
		fmt.Printf("%d-%v-%T\n",index,val,val)
	}

}