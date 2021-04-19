package main


import (
	"fmt"
	"strconv"
	"time"
)
func main() {
	startTime := time.Now().Unix()
	go test()
	go test()
	 //test()
	for i := 1;i <= 3;i++ {
		//fmt.Println("hello world" + strconv.FormatInt(int64(i),10))
		fmt.Println("main" + strconv.Itoa(i))
		time.Sleep(time.Second )
	}
	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)
}

func test() {
	for i := 1;i <= 10;i++ {
		//fmt.Println("hello world" + strconv.FormatInt(int64(i),10))
		fmt.Println("test()" + strconv.Itoa(i))
		time.Sleep(time.Second )
	}
}