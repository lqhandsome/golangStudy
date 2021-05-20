package main

import "fmt"

func main() {
	arr := [9]int{1, 2, 3, 3, 3, 3, 7, 8, 9}
	//fmt.Println(arr[(1+2)/2])
	//index := midFind(arr, 4)
	//fmt.Println(index)
	//fmt.Println(cursion(arr, 1))
	fmt.Println(FindFirstValue(arr, 3))
}

func midFind(arr [9]int, n int) (index int) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == n {
			return mid
		}
		if arr[mid] <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func cursion(arr [9]int, n int) int {
	return search(arr, 0, len(arr)-1, n)
}
func search(arr [9]int, left int, right int, val int) int {
	if left > right {
		return -1
	}
	mid := left + ((right - left) >> 1)
	if arr[mid] == val {
		return mid
	}
	if arr[mid] < val {
		left = mid + 1
		return search(arr, left, right, val)
	} else {
		right = mid - 1
		return search(arr, left, right, val)
	}

}

func FindFirstValue(arr [9]int, value int) (index int) {
	l := 0
	r := len(arr) - 1
	var mid int
	for l <= r {
		//{1, 2, 3, 3, 3, 6, 7, 8, 9}
		mid = l + ((r - l) >> 1)
		if arr[mid] < value {
			l = mid + 1
		}
		//查找第一个值等于给定值
		//if arr[mid] == value {
		//	if mid == 0 || arr[mid-1] != value {
		//		return mid
		//	}
		//	r = mid -1
		//}

		//查找最后一个值等于给定值
		if arr[mid] == value {
			if mid == len(arr) - 1 || arr[mid + 1] != value {
				 return mid
			}
			l = mid + 1
		}
		if arr[mid] > value {
			r = mid - 1
		}

	}

	return -1

}
