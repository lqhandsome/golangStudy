package main

import "fmt"


type monster struct {
	name string
	age int
	slice []int
}
func main()  {
	m1 := monster{
		"牛魔王",
		21,
		[]int{1,2,3},
	}
	var m2 *monster = &m1
	fmt.Println("m1")
	fmt.Printf("%T",m2.name)
}
