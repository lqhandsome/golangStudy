package main

import "fmt"

func main() {
	arr := [6]int{5, 1, 23, 11, 100, 89}
	//popSort(&arr)
	//selectSort(&arr)
	//insertSort(&arr)
	quickSort(&arr,0,5)
	fmt.Println(arr)

}

//冒泡排序
func popSort(arr *[6]int) {

}

//选择排序
func selectSort(arr *[6]int) {

}

//插入排序
func insertSort(arr *[6]int) {

}

//快速排序
func quickSort(arr *[6]int, l int, r int) {

}