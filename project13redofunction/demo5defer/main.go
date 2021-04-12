package main

import (
	"fmt"
)

func init()  {
	
}

func main()	{
	fmt.Println(sum(10,20))
	//defersum
}

//defer主要用于关闭资源,比如文件资源,数据库资源
// db := OpenDatabase 
//defer db.Close

//defer 讲该行代码压入栈暂不执行
//当函数要return的时候,讲defer的数据弹出
//入栈时进行值拷贝
func sum(n1 int,n2 int) int {
	defer fmt.Println("ok1 n1=",n1) // 3
	defer fmt.Println("ok2 n2=",n2) // 2

	res := n1 + n2
	fmt.Println("ok3 res=",res) // 1
	return res
}