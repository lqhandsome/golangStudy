package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main()  {
	now := time.Now()
	fmt.Println(time.Now())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Format("2006年01月-02日 15:03:04"))
	rand.Seed(now.UnixNano())
	rand.Seed(1)

	fmt.Println(rand.Intn(100))
	//获取从1970年一月一号到现在的秒数
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(time.Millisecond)
	time.Sleep(time.Millisecond * 6000)
}