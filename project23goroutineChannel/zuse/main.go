package main

import (
	"fmt"
	"time"
)

/*
阻塞会在读取的时候v := range c 或者v,ok := <-c的时候阻塞
也会在写入的时候阻塞c<-5 代码运行到这里被阻塞相当于休眠了一样
fatal error: all goroutines are asleep - deadlock 只会在遍历没有没有关闭的通道时才会报错
底层是遍历一个管道数据取完了底层编译器发现这个管道没有关闭就会报错，再取就会报错
 */
func main() {
	c := make(chan int,2)
	input(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	//for  {
	//	v,_ := <-c
	//	fmt.Println(v)
	//}
	//for v := range c {
	//	fmt.Println(v)
	//}

}

func input(c chan int){
	time.Sleep(time.Second * 3)
	c<-10
	//close(c)
}