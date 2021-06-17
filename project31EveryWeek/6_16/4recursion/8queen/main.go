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
	eightQueue(8)
	fmt.Println(num)
}

func eightQueue(row int) {
	if row == 8 {
		num++
		print(arr)
		return
	}
	for j := 0; j < 8; j++ {
		if isOk(row, j) {
			arr[row] = j
			eightQueue(row + 1)
		}
	}
}

//判断一个位置可以不可以放
func isOk(row int, clo int) bool {
	//这一列的左边和右边
	left := clo - 1
	right := clo + 1
	for i := row; i > 0; row-- {
		//上面有元素
		if arr[row -1] == clo {
			return false
		}
		//左上角
		if left >= 0 &&arr[row -1]  == left {
			return false
		}
		//右上角
		if right <= 8 && arr[row-1] == right {
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
