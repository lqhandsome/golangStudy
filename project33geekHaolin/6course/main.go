package main

import "fmt"

type a = int
var container = []string{"one","two"}
func main() {
	container := map[int]string{0:"zero",1:"one",2:"two"}
	value,ok := interface{}(container).(map[int]string)
	fmt.Println(value,ok)
	fmt.Printf("tel element is %q.\n",container[1])
}
