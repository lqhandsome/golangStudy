package main

import "fmt"

type Person struct {
	name string
	age int
	happy [2]int
	love []int
	adder *string
}
func main()  {
	//调用是如果要赋值需要全部赋值,如果单个赋值则可以
	//方式一
	var p1 Person
	fmt.Println("p1=",p1)

	//方式二
	p2 := Person{}
	fmt.Println("p2=",p2)

	//方式三
	var p3 *Person = new(Person)
	fmt.Println("p3=",p3)

	//方式四
	var p4 *Person = &Person{}
	fmt.Println("p4=",p4)

	fmt.Printf("p1p2p3p4的类型风别是%T %T %T %T\n",p1,p2,p3,p4)

	arr := [...]int{1,2,3,4}

	ptr := &(arr[0])
	fmt.Println((*ptr))
}


