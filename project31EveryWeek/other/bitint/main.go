package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "123"
	s2 := "1234"
	for i,v := range s1 {
		fmt.Println(i,v)
	}
	fmt.Println(addStrings(s1,s2))
}

func addStrings(num1 string, num2 string) string {
	len1 := len(num1) - 1
	len2 := len(num2) - 1
	resString := "" //用来保存字符串
	and := 0 //用来保存进位
	for ;len1 >= 0 || len2 >= 0; len2,len1 = len2 -1,len1-1 {
		var x,y int
		if len1 >= 0 {
			x = int(num1[len1] - 48)
		}
		if len2 >= 0 {
			y = int(num2[len2] - 48)
		}

		res := x +y +and
		resString = strconv.Itoa((res) % 10) + resString
		and = res/10
	}
	if and != 0 {
		resString = strconv.Itoa(and )+ resString
	}
	return resString
}