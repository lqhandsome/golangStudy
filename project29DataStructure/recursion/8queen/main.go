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
	arr = make([]int, 8)
	eightQueue(0)
	fmt.Println(num)
	fmt.Println(arr)
}

func eightQueue(row int) {

	//代表8行已经放完
	if row == 8 {
		print(arr)
		num++
		return
	}
	for i := 0; i < 8; i++ {
		if isOk(row, i) {
			arr[row] = i
			eightQueue(row + 1)
		}
	}
}

//判断一个位置可以不可以放
func isOk(row int, clo int) bool {
	//当前节点的左边
	left := clo - 1
	//当前节点的右边
	right := clo + 1
	for i := row - 1; i >= 0; i-- {
		//判断上一个位置是否有元素
		if arr[i] == clo {
			return false
		}
		//判断左上角是否有
		if left >= 0 && arr[i] == left {
			return false
		}
		//判断右上角是否有
		if right <= 8 && arr[i] == right {
			return false
		}
		left--
		right++
	}
	return true
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
