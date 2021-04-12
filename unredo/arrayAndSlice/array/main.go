package main

import "fmt"

func main() {

	arrOne := [6]int{1,2,3,4,5,6}
	var  arrTwo [2]int = [2]int{7,8}
	fmt.Println(arrOne,arrTwo)
	fmt.Printf("%p\n",&arrOne)
	arrSlice := arrOne[:1]
	arrSlice[0] = 8
	fmt.Println("arrSlice",arrSlice)
	fmt.Printf("%p\n",&arrOne)


	arrMap := make(map[string]string,10)
	s := make([]string,10)
	s[9] = "100"
	s = append(s,"hello")
	arrMap["lq"] = "handsome"
	fmt.Println(arrMap)
	fmt.Println(len(s))
	arrMap["lq"] = "handsome"
	s = append(s,"hello")
}
