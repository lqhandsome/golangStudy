package main

import (
	"fmt"
)
func main()  {
	// var sliceOne []int
	// fmt.Println(sliceOne)

	// arr := [4] int{1,2,3,4}
	// sliceTwo := arr[1:3]
	// sliceTwo[1] = 5
	// fmt.Println(sliceTwo)
	// fmt.Println(arr)

	slice1 := make([]float64,4)

	fmt.Println(slice1,cap(slice1))
	a := 1
	fmt.Printf("%T-%p-%p-%p-%v",slice1,&slice1,slice1,&a,*(&slice1))
}