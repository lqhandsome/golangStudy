package main

import "fmt"

func  main()  {
		//var a map[int]int
		a := make(map[int]int,5)
		a[1] = 1
		fmt.Println("a",a)
		var arr = new([6]int)
		fmt.Println("arr",arr[0])
		var sliceOne = make([]int,5)
		fmt.Println("sliceOne",sliceOne)
		var sliceTwo []int
		sliceTwo = []int{1,2,3}
		fmt.Println("sliceTwo",sliceTwo)
}
