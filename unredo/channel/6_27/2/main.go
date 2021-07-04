package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int,1),
		make(chan int,1),
	}
	//go read(ch)
	rand.Seed(int64(time.Now().UnixNano()))
	//index := rand.Intn(3)
	//fmt.Println(index)
	//intChannels[index]<- index
	intChannels[1]<-1
	intChannels[0]<-0
	intChannels[2]<-2
	select {
	case <-intChannels[0]:
		fmt.Println("the first")
	case <-intChannels[1]:
		fmt.Println("the two")
	case <-intChannels[2]:
		fmt.Println("the second")
	default:
		fmt.Println("no one")
	}
	fmt.Println(len(intChannels[0]),len(intChannels[1]),len(intChannels[2]))
}
func read(ch chan <- int) {
	for  {
		ch<-1
		time.Sleep(1)
	}
}
