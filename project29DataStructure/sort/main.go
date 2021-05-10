package main

import "fmt"

func main() {
	arr := [5]int{10,34,19,100,80}
	SelectSort(&arr)
	fmt.Println(arr)

}

func SelectSort(arr *[5]int) {
	l := len(*arr)
	for i := 0;i < l; i++ {
		for j := i +1;j < l;j++{
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
	return
}

func insertSort(arr *[5]int) {
	//l := len(*arr)
}