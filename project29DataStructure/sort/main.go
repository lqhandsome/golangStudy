package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//arr := [6]int{10, 34, 19, 100, 80, -10}
	var arr[80000]int
	for i := 0; i < 80000; i++ {
		//rand.Seed(time.Now().Unix())
		rand.Seed(time.Now().UnixNano())
		//time.Sleep(time.Second / 100)
		arr[i] = rand.Intn(9000)
	}

	start := time.Now().UnixNano()
	//SelectSort(&arr)
	//insertSort(&arr)
	quickSort(0, len(arr)-1, &arr)
	end := time.Now().UnixNano()
	fmt.Println("快速耗时",end - start)

}

func SelectSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if (arr)[i] > (arr)[j] {
				tmp := (arr)[i]
				(arr)[i] = (arr)[j]
				(arr)[j] = tmp
			}
		}
	}
	return
}

//插入排序
func insertSort(arr *[80000]int) {
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
func quickSort(left int, right int, arr *[80000]int) {
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
	arr[left] = arr[l]
	arr[l] = pivot
	quickSort(left, l-1, arr)
	quickSort(l+1, right, arr)
}

