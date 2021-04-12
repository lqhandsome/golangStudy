package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main()  {
	var arr  [6]int
	for i := 0; i < len(arr); i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second)
		arr[i] = rand.Intn(100)

	}
	fmt.Println(arr)
	count := len(arr)
	for i := 0; i < len(arr) /2 ; i++ {
		tmp := arr[i]
		arr[i] = arr[count -1 -i]
		arr[count -1 -i] = tmp
	}
	fmt.Println(arr)
}
