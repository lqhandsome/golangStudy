package main

import "fmt"

var (
	myMap = make(map[int]uint64,10)
)

func test(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = uint64(res)
	return res
}
func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	for k, v := range myMap {
		fmt.Printf("%d= %v\n",k,v)
	}
	//fmt.Println(myMap)
}
