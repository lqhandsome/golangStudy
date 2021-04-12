package main

import "fmt"

//func main() {
//	slice := []int{1,2,3,4,5}
//	fmt.Println(slice)
//	exchangeSlice(slice)
//	fmt.Println(slice)
//}
//
//func exchangeSlice(slice []int) {
//	for k, v := range slice {
//		slice[k] = v * 2
//	}
//}

func sliceModify(slice []int) []int{
	// slice[0] = 88
	slice = append(slice, 6)
	return  slice
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = sliceModify(slice)
	fmt.Println(slice)
}

