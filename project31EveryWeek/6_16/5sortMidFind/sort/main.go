package main

import "fmt"

func main() {
	arr := [6]int{5, 1, 23, 11, 100, 89}
	//popSort(&arr)
	//selectSort(&arr)
	//insertSort(&arr)
	quickSort(&arr, 0, 5)
	fmt.Println(arr)

}

//冒泡排序
func popSort(arr *[6]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//选择排序
func selectSort(arr *[6]int) {
	for i := 0; i < len(arr); i++ {
		//tmp := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//插入排序
func insertSort(arr *[6]int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for ; j > 0 && tmp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

//快速排序
func quickSort(arr *[6]int, l int, r int) {
	left := l
	right := r

	if left >= right {
		return
	}
	pivot := arr[l]
	for left < right {
		for right > left && arr[right] >= pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		//把数据交换
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	//把主元换过来
	arr[l] = arr[left]
	arr[left] = pivot
	quickSort(arr, l, left-1)
	quickSort(arr, right+1, r)

}
