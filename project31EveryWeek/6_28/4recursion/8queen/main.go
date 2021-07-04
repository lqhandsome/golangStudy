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
	arr = make([]int,8)
	eightQueue(0)
	fmt.Println(num)
}

func eightQueue(row int) {
	if row == 8 {
		num++
		print(arr)
		return
	}
	for i := 0;i < 8 ;i++ {
		if isOk(row,i)  {
			arr[row] = i
			eightQueue(row + 1)
		}
	}
}

//判断一个位置可以不可以放
func isOk(row int, clo int) bool {
	left,right := clo -1,clo + 1
	for i := row -1;i >= 0;i-- {
		//判断上面
		if arr[i] == clo {
			return false
		}
		//判断左上角
		if left >= 0 && arr[i] == left {
			return false
		}
		//判断右上角
		//判断左上角
		if right <=7 && arr[i] == right {
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
