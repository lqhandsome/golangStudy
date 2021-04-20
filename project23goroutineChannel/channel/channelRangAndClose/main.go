package main

import (
	"fmt"
	"time"
)

//向管道写入五十个数据后关闭
func writeData(c chan int) {
	for i := 1;i <= 50;i++{
		c<-i
		fmt.Println("写入数据=",i)
		time.Sleep(time.Second)
	}
	close(c)
}

//读取数据
func readData(c chan int,exitChan chan bool) {
	for  {
		v,i := <-c
		if !i {
			break
		}
		fmt.Println("读取到数据=",v)

	}
	close(exitChan)
}

func main() {

	intData := make(chan int,10)
	exitData := make(chan bool,1)
	go  writeData(intData)
	go readData(intData,exitData)

	for  {
		_,ok := <-exitData
		if !ok {
			break
		}
	}

}
