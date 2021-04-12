package main

import (
	"fmt"
	"math"
)

func main()  {
//单分支
	if 1 < 2 {
		fmt.Println(111)
	} else {
		fmt.Println(222)
	}
	//可以这样定义 i只能在这里使用
	if i := true;i {
		fmt.Println("true")
	}
//多分支
	if 1 == 2 {
		fmt.Println(30)
	} else if 1 < 2 {
		fmt.Println(32)
	} else {
		fmt.Println(34)
	}

	var score int
	fmt.Scanln(&score)	
	// fmt.Scanf("%d",&score)	

	if score == 100 {
		fmt.Println("奖励BWM")
	} else if score > 90 {
		fmt.Println("奖励iphone")
	} else if score > 80 {
		fmt.Println("奖励IPAD")
	} else  {
		fmt.Println("什么都不奖励")
	}
	var a , b , c float64 = 3 ,1 ,60

	fmt.Scanf("%f %f %f",&a,&b,&c)
	m :=b * b - 4 * a * c
	//多分支判断

	if m < 0 {
		fmt.Println("此方程无解")
	} else if m  > 0{
		x1 := (-b + math.Sqrt(m))
		x2 := (-b - math.Sqrt(m))
		fmt.Println("x1",x1)
		fmt.Println("x2",x2)
		} else {
			x2 := (-b - math.Sqrt(m))
			fmt.Println("x1",x2)
			fmt.Println("一个解")
		}
		if ii :=10; ii<10{
				fmt.Println(111)
		}	
}