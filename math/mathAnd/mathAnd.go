package main

import (
	"fmt"
)
func main() {
	 i := 11
	fmt.Printf("%T,%v",10.5 / 4,10.5 / 4)
	i++
	fmt.Printf("%v\r",float64(10) / 4)

	var lastDays int = 75
	var week int = lastDays / 7
	var day  int = lastDays % 7
	fmt.Printf("%v%v\r",week,day)

	var 花式 float32 = 134.5

	var 摄氏  = 5.0 / 9 * ( 花式 - 100 )
	fmt.Printf("%v",摄氏)

	t := i -1
	fmt.Printf("%v,%T",t,t)
	var a, b  = 10, 9
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a,b)

	a, b  = 100, 90
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a,b)
}