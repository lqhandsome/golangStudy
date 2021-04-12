package main

import (
	"fmt"
)

func  main()  {
	var arr [6]int = [...]int{11,22,33,44,55,66}
	slice1 := arr[:]
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	for i,v := range slice1 {
		fmt.Println(i,v)
	}
}