package main

import "fmt"

func main() {
	arr := [6]int{10, 34, 19, 100, 80, -10}
	//SelectSort(&arr)
	//insertSort(&arr)
	quickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)

}

func SelectSort(arr *[5]int) {
	l := len(*arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
	return
}

//插入排序
func insertSort(arr *[5]int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		//要插入的数据
		tmp := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

//快速排序 {10, 34, 19, 100, 80,-10}
func quickSort(left int, right int, arr *[6]int) {
	//定义左边界
	l := left
	//定义右边届
	r := right
	if l > r {
		return
	}
	//主元
	pivot := arr[l]

	for l != r {
		//从右边开始
		for arr[r] >= pivot && l < r {
			r--
		}

		//从左边开始
		for arr[l] <= pivot && l < r {
			l++
		}

		//符合这交换
		if l < r {
			arr[r], arr[l] = arr[l], arr[r]
		}
	}
	//将主元换到中间
	arr[left] =  arr[l]
	arr[l] = pivot
	quickSort(left,l -1,arr)
	quickSort(l +1 ,right,arr)
}
