package main

import "fmt"

type ValNode struct {
		row int
		col int
		val interface{}
}
func main() {
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //篮子

	//输出查看数据
	for _,v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组的规模（行和列，默认值）
	valNode := ValNode{
		11,
		11,
		0,
	}
	sparseArr = append(sparseArr,valNode)
	for i,v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个子节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	fmt.Println(valNode)
	fmt.Println(sparseArr)
}
