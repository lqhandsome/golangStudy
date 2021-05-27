package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arrN(arr, 6, 6)

}

func arrN(arr [6]int, len int, k int) {
	if k == 1 {
		for i := 0; i < len; i++ {
			fmt.Print(arr[i])
		}
		fmt.Println()
	}
	for i := 0; i < k; i++ {
		arr[i], arr[k-1] = arr[k-1], arr[i]
		arrN(arr, len, k-1)
		arr[i], arr[k-1] = arr[k-1], arr[i]
	}

}
