package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

var (
	arr = [8]int{}
	num int
)
//实现一个环形链表
func main() {
	cal8queens(0)
	fmt.Println(num)
}
/**
	row放到了多少行
column 列
row 行
*/
func cal8queens(row int) {
	if row == 8 { //代表8行已经全部放完
		//打印出一个数组
		num++
		return
	}
	for column := 0 ;column < 8 ;column++{
		if isOk(row,column) {
			//fmt.Println(1)
			arr[row] = column
			cal8queens(row+1)
		}
	}
}

//用来检测是否能放节点
func isOk(row int,column int) bool {
	//判断该位置的左边和右边
	left := column - 1
	right := column +1
	//从上一行开始判断是否可以放入
	for i := row -1;i >= 0;i-- {

		//如果上面已经有点了
		if arr[i] == column {
			return false
		}

		//考察左上角是否有
		if left > 0  {
			if arr[i] == left {
				return  false
			}
		}
		//考察右上角是否有
		if right < 8  {
			if arr[i] == right {
				return  false
			}
		}
		left--
		right++
	}
	//如果发现可以放
	return true
}