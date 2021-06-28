package main

import (
	"fmt"
	"time"
)

func main() {
	chanOne  := make(chan interface{},2)
	//go pop(chanOne)
	go pop(chanOne)
	//close(chanOne)
	chanOne<-1
	//time.Sleep(time.Second * 3)

	for v := range chanOne{
		fmt.Println(v)

	}
}

func pop(c chan interface{}) {
	for true {
		time.Sleep(time.Second * 1)
		c<-1
	}
}