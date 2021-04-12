package main

import (
	"fmt"
	"unsafe"
)
/*

*/
func main()  {
	var price  int16= '刘'
	var bu bool = true

	fmt.Printf("%c %T %d\n",price,price,unsafe.Sizeof(price))
	fmt.Printf(" %T %d",bu,unsafe.Sizeof(bu))
}