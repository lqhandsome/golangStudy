package main

import (
	"fmt"
)

func main()  {
	var byteArr [26] byte
	
	byteArr[0] = 'A'
	fmt.Println(byteArr[0])
	
	fmt.Println(byteArr)
	for index,_ := range byteArr {
		if index == 0 {
			continue
		}
		if index >= len(byteArr) -1{
			break
		}
		byteArr[index] = byteArr[index -1] + 1
		fmt.Printf("%c ",byteArr[index])
	}
	var i byte = 0;
	for ; i < byte(len(byteArr)); i++ {
		byteArr[i] = 'A' + i
	}
	fmt.Printf("%c",byteArr)
	fmt.Printf("%c",100)
}