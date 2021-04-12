package main

import (
	"fmt"
	"errors"
)

func main(){
    defer func(){
        if err := recover() ; err != nil {
            fmt.Println(err) //three
        }
    }()
    defer func(){
        panic("three")
    }()
    defer func(){
        panic("two")
    }()
    panic("one")
}
// func main()  {
// 	//异常处理
// 	// test()
// 	// fmt.Println("下面的程序继续执行")
// 	defer eee()
// 	paincFun()
// }

func paincFun(){
	err := throwErr()
	if err != nil {
		fmt.Printf("%T\n",err)
		panic(err)
	}
	fmt.Println(err)
}

func throwErr() (error) {
	if true {
		return errors.New("16行的错误")
	} else {
		return nil
	}
}

func test(){
	defer eee()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func eee()  {
	err := recover()
	if err != nil {
		fmt.Println("err=",err)
	}
}