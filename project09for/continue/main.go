package main

import (
	"fmt"
)
func main()  {
	for i := 0; i < 10; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	label1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				continue label1
			}
			fmt.Println(i,"_",j)
		}
	}
}