package main

import (
	"fmt"
	// "strings"
	// "strconv"
)

var (
	fun = func(n1 int ,n2 int)int {
		return n1 + n2
	}
)

type myFunctionType func(i int,ii int)int
func main()  {
	//自定义变量类型
	type bytes byte
	var bb bytes = 'c'
	fmt.Println(bb)
	fmt.Printf("%T\n",bb)
	
	res := test(myFunction,1,2)
	fmt.Printf("%T%v\n",res,res)

	fmt.Printf("%T%v\n",changeParms(1,2,3,4),changeParms(1,2,3,4))
	fmt.Printf("匿名函数fun的类型是%T基本值是%v\n",fun(10,20),fun(10,20))
	func(){
		fmt.Println("29行的匿名函数",bb)
	}()
	res2 := changeParms
	fmt.Println(res2(10,20,30))
}

func test(myfunctionParm myFunctionType,n1 int,n2 int)int {
		return myfunctionParm(n1,n2)
}

func myFunction(i int,ii int)int{
	return i + ii
}
func changeParms(n1 int,args... int)(sum int){
	sum = n1
	for index,_ := range args{
		sum += args[index]
	}
	return 
}