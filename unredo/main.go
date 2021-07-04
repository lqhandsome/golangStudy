package main

import "fmt"

func main() {
	num1,num2 := 111,899
	for num2 != 0 {
		tmp := num1
		num1 = num2 ^ num1
		num2 = (tmp & num2) << 1
	}
	fmt.Println(num1)
}

func modify(args [3][]string) {
	//args[1] = []string{"1","2","3"}
	args[1][1] = "1"
}