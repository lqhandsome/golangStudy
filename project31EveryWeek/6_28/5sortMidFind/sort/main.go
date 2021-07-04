package main

import "fmt"

func main() {
	arr := [7]int{5, 1, 23, 11, 100, 89,1}
	//popSort(&arr)
	//selectSort(&arr)
	//insertSort(&arr)
	quickSort(&arr,0,6)
	fmt.Println(arr)

}

//冒泡排序
func popSort(arr *[6]int) {
	for i := 0;i < len(arr);i++ {
		for j := 1;j < len(arr) - i;j++ {
			if arr[j] < arr[j -1] {
				arr[j],arr[j - 1] = arr[j - 1],arr[j]
			}
		}
	}
}

//选择排序
func selectSort(arr *[6]int) {
	for i := 0;i < len(arr);i++ {
		for j := i+1;j < len(arr);j++ {
			if arr[j] < arr[i] {
				arr[j],arr[i] = arr[i],arr[j]
			}
		}
	}
}

//插入排序
func insertSort(arr *[6]int) {
	for i := 1;i < len(arr);i++ {
		tmp := arr[i]
		j := i
		for ;j > 0 && arr[j-1] > tmp;j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

//快速排序
func quickSort(arr *[7]int, l int, r int) {
	left,right := l,r
	if l >= r {
		return
	}
	//选主元
	pivot := arr[l]

	for left < right {
			//从右边开始
		for left < right && arr[right] >= pivot {
			right--
		}
		//遍历右边
		for left < right && arr[left] <= pivot {
			left++
		}
		//一次遍历完，交换
		if left < right {
			arr[left] , arr[right] = arr[right],arr[left]
		}
	}
	//交换玩把主元换到中间
	arr[l] = arr[left]
	arr[left] = pivot
	fmt.Println(arr) //5, 1, 23, 11, 100, 89
	//return
	quickSort(arr,l,left-1)
	quickSort(arr,right+1,r)
}