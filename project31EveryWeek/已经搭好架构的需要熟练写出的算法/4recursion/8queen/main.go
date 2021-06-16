package main

import (
	"fmt"
)

var (
	//原来统计有多少个
	num int
	//记录数据
	arr []int
)

func main() {

}

func eightQueue(row int) {

}

//判断一个位置可以不可以放
func isOk(row int, clo int) bool {

}
func print(arr []int) {
	for _, v := range arr {
		for i := 0; i < 8; i++ {
			if i == v {
				fmt.Print("* ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Println("----------------------------")
}
