package main

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

type T struct {
	n  string
	i  int
	f  float64
	fd *os.File
	b  []byte
	s  bool
}
func main() {

	var t1 *T
	t1 = new(T)
	fmt.Println(t1)

	t2 := T{}
	fmt.Println(t2)

	t3 := T{"hello", 1, 3.1415926, nil, make([]byte, 2), true}
	fmt.Println(t3)
	return

	var wg sync.WaitGroup
	ch := make(chan int,64)
	sy := make(chan int,2)
	wg.Add(1)
	go Producer(3,ch,&wg,sy)
	wg.Add(1)
	go Producer(5,ch,&wg,sy)
	go Consumer(ch,sy)
	//time.Sleep(5 * time.Second)
	close(ch)
	wg.Wait()

}

//生产者：生成factor整数倍的序列
func Producer(factor int,out chan<- int,wg *sync.WaitGroup,sy chan int) {
	for i:=0;;i++{
		out <- factor * i
		if i > 10 {
			break
		}
	}
	sy<-1
	 wg.Done()
}

//消费者
func Consumer(in <-chan int,sy chan int) {
	for v := range in {
		if cap(sy) == 2 {
			break
		}
		fmt.Println(v)
	}
}
func TestAlloc(t *testing.T) {
	type T struct {
		n  string
		i  int
		f  float64
		fd *os.File
		b  []byte
		s  bool
	}

}