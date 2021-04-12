package main

import (
	"fmt"
)
func main()  {
	//break退出循环
	// break加标签可以跳出到指定循环
	// label1:
	for i := 0; i < 5; i++ {
		label2:
		for j := 0; j < 4; j++ {
			if(j == 1) {
				break label2;
			}
			fmt.Println(i)
		}
	}
}