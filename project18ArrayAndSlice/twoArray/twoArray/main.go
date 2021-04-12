package main

import (
	"fmt"
)

func main()  {
	var twoArray [2][2]int
	arr := new([2][2][2]int)
	arr[1][1][1] = 11
	crr := [...][2]int{{1,2},{2,3}}
	fmt.Println("二维数组")
	fmt.Println(twoArray)
	fmt.Println(*arr)
	fmt.Println("crr等于",crr)
}