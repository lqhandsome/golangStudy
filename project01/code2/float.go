package main

import (
	"fmt"
	"unsafe"
)
/*

*/
func main()  {
	var price float32 = 5.1234e-3
	fmt.Println(price)
	fmt.Printf("%T %d",price,unsafe.Sizeof(price))
}