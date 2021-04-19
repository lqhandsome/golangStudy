package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]uint64,10)
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = uint64(res)
	lock.Unlock()
	//fmt.Printf("myMap[%d]=%d\n",n,res)
}
func main() {
	//执行多个任务200个
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠十秒等待协程执行结束
	time.Sleep(time.Second * 3)

	//遍历打印myMap
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("%d= %v\n",k,v)
	}
	lock.Unlock()
}
