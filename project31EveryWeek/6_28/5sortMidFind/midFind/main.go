package main

import "fmt"

func main() {
	arr := [9]int{1, 2, 3, 3, 3, 3, 7, 8, 9}
	fmt.Println(findLastVal(arr,3))
}

func midFind(arr [9]int,val int) (index int) {
	var mid int
	var l int
	var r = len(arr)
	mid = (0 + len(arr)) >> 1
	for l <= r {
		if arr[mid] == val {
			return mid
		}
		if arr[mid] < val {
			l = mid +1
			mid = (l + r) >> 1
		}
		if arr[mid] > val {
			r = mid -1
			mid = (l + r) >> 1
		}
	}
	return  -1
}

func findFirstVal(arr [9]int,val int) (index int) {
	var mid int
	var l int
	var r = len(arr)
	mid = (0 + len(arr)) >> 1
	for l <= r {
		if arr[mid] == val {
			//第一个等于
			for mid >=0 && arr[mid] == val {
				mid--
			}
			return mid + 1
		}
		if arr[mid] < val {
			l = mid +1
			mid = (l + r) >> 1
		}
		if arr[mid] > val {
			r = mid -1
			mid = (l + r) >> 1
		}
	}
	return  -1
}

func findLastVal(arr [9]int,val int) (index int) {
	var mid int
	var l int
	var r = len(arr)
	mid = (0 + len(arr)) >> 1
	for l <= r {
		if arr[mid] == val {
			//第一个等于
			for mid <= 8 && arr[mid] == val {
				mid++
			}
			return mid - 1
		}
		if arr[mid] < val {
			l = mid +1
			mid = (l + r) >> 1
		}
		if arr[mid] > val {
			r = mid -1
			mid = (l + r) >> 1
		}
	}
	return  -1
	return
}
