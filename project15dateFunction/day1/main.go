package main

import (
	"fmt"
	"time"
)
const (
	Hour = 50 * 20
)

func main(){
	fmt.Println("dateFunction")
	fmt.Printf("%T %v\n",time.Now(),time.Now())
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Day())
	//格式化时间方式1
	strDate := fmt.Sprintf("%v-%d-%d-%d",now.Year(),int(now.Month()),now.Hour(),now.Minute())
	fmt.Println(strDate)

	//格式化时间方式2
	fmt.Println(now.Format("2006-01-02 15:04:06"))
	i := 0
	for {
		i++
		fmt.Println(i)
		fmt.Println(time.Millisecond * 100)
		time.Sleep(time.Millisecond * 100)
		if i == 10 {
			break
		}
	}
}