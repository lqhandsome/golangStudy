package main

import "fmt"

func main() {
	//创建二维数组模拟迷宫
	//1.y元素的值为1代表墙，元素为0代表以前没走过 元素为2是一个通路 元素值为3代表走过但是走不通
	var arr [8][7]int

	//上下值设为1
	for i := 0; i < 7; i++ {
		arr[0][i] = 1
		arr[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		arr[i][0] = 1
		arr[i][6] = 1
	}
	arr[3][1] = 1
	arr[3][2] = 1
	SetWay(&arr,1,1)
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func SetWay(arr *[8][7]int, i int, j int) bool {
	if arr[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if arr[i][j] == 0 {
			arr[i][j] = 2
			if SetWay(arr, i+1, j) { //下
				return true
			} else if SetWay(arr, i, j+1) { //右
				return true
			} else if  SetWay(arr, i -1 , j ) { //上
				return true
			} else if SetWay(arr, i, j - 1) { //左
				return true
			} else {
				//是条死路
				arr[i][j] = 3
				return false
			}

		}  else {
			return false
		}
	}
}
func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println(n)
	}
}
