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
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

//选择排序
func selectSort(arr *[6]int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

//插入排序
func insertSort(arr *[6]int) {
	for i := 1;i < len(arr);i++ {
		tmp := arr[i]
		j := i
		for ; j > 0 && arr[j -1] > tmp;j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
		
	}
}

//快速排序
func quickSort(arr *[6]int,l int,r int) {
	low := l
	high := r
	if l > r {
		return
	}
	pivot := arr[low]

	for  low <= high {
		for arr[high] >= pivot && low < high {
			high--
		}
		for arr[low] <= pivot && low < high {
			low++
		}
		arr[l]= arr[low]
		arr[low] = pivot
		quickSort(arr,low,l-1)
		quickSort(arr,high+1,r)
	}

}
