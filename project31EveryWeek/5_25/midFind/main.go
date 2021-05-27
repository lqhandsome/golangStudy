package main

import "fmt"

func main() {
	arr := [9]int{1, 2, 3, 3, 3, 3, 7, 8, 9}
	fmt.Println(midFind(arr,3))
	fmt.Println(findFirstVal(arr,3))
	fmt.Println(findLastVal(arr,3))
}

func midFind(arr [9]int,val int) (index int) {
	l := 0
	r := len(arr)
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == val {
			return mid
		}
		if val < arr[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return  -1
}

func findFirstVal(arr [9]int,val int) (index int) {
	l := 0
	r := len(arr)
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == val {
			for arr[mid] == val {
				mid--
			}
			return  mid +1

		}
		if val < arr[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return  -1
}

func findLastVal(arr [9]int,val int) (index int) {
	l := 0
	r := len(arr)
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == val {
			for arr[mid] == val {
				mid++
			}
			return  mid - 1

		}
		if val < arr[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return  -1
}
