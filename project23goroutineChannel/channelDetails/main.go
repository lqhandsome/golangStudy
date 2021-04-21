package main

import "fmt"

//select可以解决从管道取数据的阻塞问题

func main() {
	intChan := make(chan int,10)
	strChan := make(chan string,10)
	for i := 0;i < 10;i++ {
		intChan<- i+1
		strChan<- fmt.Sprintf("hello world-%d",i)
	}
	label1:
	for {
		select {
			case v := <-intChan:
				fmt.Println(v)
			case v := <-strChan :
				fmt.Println(v)
		default:
			fmt.Println("没有任何数据了")
			break label1
		}

	}
}