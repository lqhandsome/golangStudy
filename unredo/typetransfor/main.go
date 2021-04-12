package main

import "fmt"
import _ "strconv"
var (
	 intOne int = 1
	 intTwo bool = true
	 intThree float32 = 3.15
	 intFour byte = 'c' 

)
// func main()  {
// 	var s1 =  fmt.Sprintf("%d",intOne) 
// 	fmt.Printf("类型是%T,值是%v\r",s1,s1)

// 	s1 = fmt.Sprintf("%t",intTwo)
// 	fmt.Printf("类型是%T,值是%v\r",s1,s1)

// 	s1 = fmt.Sprintf("%f",intThree)
// 	fmt.Printf("类型是%T,值是%v\r",s1,s1)

// 	s1 = fmt.Sprintf("%c",intFour)
// 	fmt.Printf("类型是%T,值是%v\r",s1,s1)

// }

func  main()  {
	ii := 5
	res := addr(&ii)
	fmt.Println(ii)
	fmt.Println(res)
	str := "liuqiang"
	for index,value := range str {
		fmt.Printf("%v%c\n",index,value)
	}
}

func addr(i *int) int {
	*i = *i +1
	fmt.Println(i)
	return *i
}