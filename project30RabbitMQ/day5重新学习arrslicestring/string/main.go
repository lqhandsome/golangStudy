package main

import "fmt"

func main() {
	var a [3]int
	var b = [...]int{1,2,3}
	var c = [...]int{1:2}
	aa := &a
	fmt.Println(a,b,c,cap(a),aa)
	for index,val := range aa{
		fmt.Println(index,val)
	}
}
