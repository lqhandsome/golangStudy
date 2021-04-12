package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(111)

	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:03:04"))
	fmt.Println(now.Weekday())
	fmt.Println(now.Location())
	timea := time.ParseError{"12121","212213","1212","123231","你好"}

	fmt.Println(timea.Error())
	fmt.Println(now.UnixNano())
}
