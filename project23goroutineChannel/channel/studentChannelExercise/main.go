package main

import "fmt"
var (
	count = 20
	i = 0
)

func writeChannel(c chan int,isWrite chan  int){
	for i := 1;i <= count; i++{
		c<-i
	}
	close(c)
	close(isWrite)
}

func sumChannel(c chan int,r chan int){
	for  {
		if i == 20 {
			close(r)
			break
		}
		v,_ := <-c
		res := 0
		for i := 1;i <= v;i++ {
			res += i
		}
		fmt.Println(i)
		r<-res
		i++
		fmt.Println("i=",i)

	}
}

//启动20个协程计算1-20的累加
func main() {

	//定义储存2000个数的channel
	numChan := make(chan int,count)

	//定义储存计算的值channel
	resChan := make(chan int,count)

	//定义一个是否计算完成的channel
	isWrite := make(chan int,1)
	go writeChannel(numChan,isWrite)

	for  {
		_,ok := <-isWrite
		if !ok {
			break
		}
	}
	fmt.Println(len(numChan))
	for i := 0; i < 8; i++ {
		go sumChannel(numChan,resChan)
	}



	fmt.Println(len(resChan))
	i := 1
	for v := range resChan {

		fmt.Printf("res[%v]=%v\n",i,v)
		i++
	}
}
