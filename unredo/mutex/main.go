package main

import (
	"fmt"
	"sync"
)

var (
	lock sync.Mutex
	m = make(map[int]interface{})
)

func main() {
	var count = 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
