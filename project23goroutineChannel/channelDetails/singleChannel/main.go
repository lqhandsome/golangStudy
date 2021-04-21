package main

import "fmt"

func main() {

	//申明管道默认双向
	var ch chan int
	ch = make(chan int,1)

	//申明只读管道
	var ch2 <-chan int
	//申明只写管道
	var cha3 chan<- int
	fmt.Println(cap(ch),ch2,cha3)
	//z作用：一般用来给函数传参使用
	write(ch)
	read(ch)
}

func write(ch chan<- int ){

}
func read(ch <-chan int ){

}



