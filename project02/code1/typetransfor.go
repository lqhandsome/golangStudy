package main

import (
	"fmt"
	"strconv"
)
//基本类型转换 
func main()  {
	var i int = 54
	var f1 float64 = 3.1415926
	var b bool = true
	var myChar byte = 'a'
	var str string = "aaa"

	str = fmt.Sprintf("%d",i)
	fmt.Printf("%T%v\r",str,str)
	
	str = fmt.Sprintf("%f",f1)
	fmt.Printf("%T%v\r",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("%T%v\r",str,str)

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("%T%v\r",str,str)

	var (
		int1 int = 99
		f2 float64 = 3.1415926
		b2 bool = false
	)

	str = strconv.FormatInt(int64(int1),10)
	fmt.Printf("%v\r",str)

	str = strconv.FormatFloat(f2,'f',10,64)
	fmt.Printf("%v\r",str)

	str = strconv.FormatBool(b2)
	fmt.Printf("%v",str)
}