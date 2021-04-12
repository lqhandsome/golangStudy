package main

import (
	"fmt"
	
)

func main()  {
	var (
		
		name string
		age int
		order float64
		isPass bool
	)

	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水")
	fmt.Scanln(&order)

	fmt.Println("是否通过了考试")
	fmt.Scanln(&isPass)
	// fmt.Scanf("%s %d %f %t",&name,&age,&order,&isPass)
	
	fmt.Printf("%T\n姓名=%v\n年龄%v\n薪水%v\n是否通过考试%v",name,name,age,order,isPass)
}