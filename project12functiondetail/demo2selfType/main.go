package main

import (
	"fmt"
)

type myFunType func(int,int) int

func myFun(myF myFunType,n1 int ,n2 int) int{
	return myF(n1 , n2)
}
func getSum(n1 int,n2 int) int{
	sum := n1 +n2
	return sum
}
//支持对函数返回值命名
func myFun2(n1 int,n2 int) (sum int,sub int){
	sum = n1 + n2
	sub = n1 - n2
	return 
}
func manyArgs(n1 int,args...int)(sum int){
	sum = n1
	for _,value := range args{
			sum += value
	}
	for i := 0;i < len(args);i++{
		sum += args[i]
	}
	return 
}

func main()  {
	
	//自定义数据类型,语法 type selfType  type
	type myInt int
	var my myInt = 40
	var ii int
	ii = int(my)
	fmt.Println(ii)
	fmt.Println(myFun(getSum,1,2))
	fmt.Println(manyArgs(1,2,3,4,5))

}
