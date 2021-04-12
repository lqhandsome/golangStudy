package main

import(
	"fmt"
	"errors"
)

func main()  {
	// defer func(){
	// 	err := recover()
	// 	if err != nil {
	// 		fmt.Println("代码有误",err)
	// 	}
	// }()
	// res := motolThrowError()
	// if res != nil {
	// 	fmt.Println(res)
	// }
	var arr  [6] float64
	arr[0] = 1.2

	num := 0.0
	fmt.Printf("%T",num)
	num = 10.01
	fmt.Printf("%T",num)
	fmt.Println(111)
	
}


func throwError()int{
	num1 , num2 := 1,0
	num3 := num1 / num2
	return num3
}

func motolThrowError()error{
	// err := throwError()

	if true {
		panic("runtime error")
	}
	return errors.New("mei错误")
}