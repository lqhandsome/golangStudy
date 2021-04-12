package main

import (
	"fmt"
)
func main()  {
	// var sum int = 0
	// var count int =0
	// for i := 1; i < 100; i++ {
	// 	if(i % 9 == 0) {
	// 		count++
	// 		sum += i
	// 	}
	// }
	// fmt.Println(sum,count)
	//打印空心金字塔
	var floor int
	fmt.Scanf("%d",&floor)
	for i := 1; i <= floor; i++ {
		 //打印空格
		for kg:=0;(floor - i) > kg;kg++{
			fmt.Print(" ")
		 }
		 //打印*带空心的
		for j := 0; j < 2*i -1 ; j++ {			
			if j ==0 || j == 2*i -2 || i == floor {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}