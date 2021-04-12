package main

import (
	"fmt"
	_ "time"
)

func main()  {
	// arr := [6]int{11,22,33,44,55,66}
	// fmt.Printf("arr[1]的地址是%p\n",&(arr[2]))
	// sliceOne := arr[1:3]
	// fmt.Println(sliceOne)
	

	// startTime := time.Now().UnixNano()
	// sliceOne = append(sliceOne,sliceOne...)
	// fmt.Println(sliceOne)
	// fmt.Println(arr)
	// fmt.Printf("arr[1]的地址是%p\n",&arr[2])
	// endTime := time.Now().UnixNano()

	// fmt.Println(float64(endTime -startTime)/1000 /1000)
	// fmt.Println(ff(1,2,3,4))

	slice1 := make([]int,5)
	slice1 = append(slice1,10)
	fmt.Println(slice1)
	fmt.Println(cap(slice1)) //5
	// rangeString()
}

func ff(ii int,args ...int)int{
	for i := 0; i < len(args); i++ {
		ii += args[i]
	}

	return ii
}	

func rangeString(){
	str := "hello world"
	fmt.Println("str的原始字节长度是",len(str))
	stringSlice := []rune(str)
	fmt.Println("stringSlice原始字节长度是",len(stringSlice))
	stringSlice[0] = '刘'
	str = string(stringSlice)
	fmt.Println(str)
	fmt.Println("str的后来字节长度是",len(str))
	fmt.Println("stringSlice原始字节长度是",len(stringSlice))
	fmt.Printf("%c",stringSlice)
	fmt.Printf("%c",str[0])
}