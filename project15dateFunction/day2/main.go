package main

import (
	"strconv"
	"fmt"
	"time"
)

func test()  {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}

}

func main()  {
	startTime := time.Now().UnixNano()
	test()
	
	time.Sleep(time.Second)
	endTime := time.Now().UnixNano()
	fmt.Println((float64(endTime - startTime ) / 1000 /1000 /100))
}