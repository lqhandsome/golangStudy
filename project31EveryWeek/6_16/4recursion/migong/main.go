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

}

func SetWay(arr *[8][7]int, i int, j int) bool {

}
func test(n int) {

}
