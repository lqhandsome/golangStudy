package main

import "fmt"

func main()  {
	fmt.Print(test())
	// label:
	// 	fmt.Print(1)
	// goto label
}

func test() int {
	return 111
}