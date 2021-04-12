package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main()  {
	var arr1 [6]int = [...] int{1,2,3,4,5,6}
	var arr2  = [...] int{1,2,3,4,5,6,7}
	arr3 := [...] int {1,2,3}
	var arr [2]int
	arr[0] = 1
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr)
	// rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}