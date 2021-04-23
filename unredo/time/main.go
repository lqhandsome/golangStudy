package main

import (
	"fmt"
	"time"
)
func main() {
	date := time.Now()

	fmt.Println(date.Format("2006-01月-02日 15时:04分:05秒"))
	fmt.Println(time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")) //2014-01-07 09:32:12
	//时间转为时间戳
	fmt.Println(time.Date(2021, 4, 15, 15, 3, 30, 0, time.Local).Unix())
	fmt.Println(date.Unix())
	fmt.Println(time.Now())
}

