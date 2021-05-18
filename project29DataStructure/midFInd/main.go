package main

import "fmt"

func main() {
	arr := [9]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr[(1+2)/2])
	index := midFind(arr,4)
	fmt.Println(index)
	fmt.Println(cursion(arr,1))
}

func midFind(arr [9]int, n int) (index int) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) /2
		if arr[mid] == n {
			return mid
		}
		if arr[mid] <= n{
			left = mid  + 1
		} else {
			right = mid - 1
		}
	}
	return  -1
}
func cursion (arr [9]int, n int) int {
	return search(arr,0,len(arr) - 1 ,n)
}

func search(arr [9]int, left int,right int,val int) int {
	if left > right {
		 return  -1
	}
	mid := left + ((right - left) >> 1)
	if arr[mid] == val {
		return mid
	}
	if arr[mid] < val {
		left = mid + 1
		return search(arr,left,right,val)
	} else {
		right = mid - 1
		return search(arr,left,right,val)
	}

}
