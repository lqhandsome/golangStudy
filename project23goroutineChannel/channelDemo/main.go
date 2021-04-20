package main


import (
	"fmt"
)

func  main() {
	var intChan chan int

	intChan = make(chan int, 3)
	intChan <- 3
	intChan <- 3
	print(intChan, "\n")
	fmt.Println(len(intChan))
	fmt.Println(cap(intChan))
	m := make(map[string]int, 3)
	fmt.Println(len(m))
}