package main

import "time"

func main() {
	//m := make(map[int]int)
	var m = make(map[int]int,10)
	//var wg sync.WaitGroup
	//wg.Add(2)
	go func() {
		for  {
			_ = m[2]
			//m[1] = 1
		}
	}()
	go func() {
		for  {
			_ = m[2]
		}
	}()
	time.Sleep(time.Second)
	//wg.Wait()
}
