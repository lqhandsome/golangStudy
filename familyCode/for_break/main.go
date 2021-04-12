package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var count int
		for {
		rand.Seed(time.Now().UnixNano())
		m := rand.Intn(100)
		count++
		if m == 99{
			fmt.Println(count)
			break
		}
	}

	fmt.Println(time.Now().Unix())
}