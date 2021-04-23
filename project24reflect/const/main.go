package main

import "fmt"

const (
	a = iota
	b
	c
	d
)
func main() {
	const (
		name = "a"
	)
	fmt.Println(name)
	test()
}

func test(){
	fmt.Println(a)
}