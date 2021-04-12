package main

import (
	"fmt"
)
func main()  {
	intArr := [6] int64{1,2,3,4,5,6}
	fmt.Printf("%p\n",&intArr) //0xc00000c440 数组的首地址 第一个元素的地址
	p := &intArr[0]
	fmt.Printf("%p-%v-%T",p,p,p) 

	var scnfArr [6]float64
	for i := 0; i < len(scnfArr) -5; i++ {
		fmt.Println("请输入第",i +1,"个元素的值")
		fmt.Scanln(&scnfArr[i])
	}
	fmt.Println(scnfArr)
	//定义数组的几个方式
	//1
	var arr1 [3]int = [3]int{1,2,3}
	//2
	var arr2 = [3]int{3,4,5}
	//3
	var arr3 = [...]int{5,6,7}
	//4
	arr4 := [1]string{"1"}
	fmt.Println(arr1,arr2,arr3,arr4)
}
