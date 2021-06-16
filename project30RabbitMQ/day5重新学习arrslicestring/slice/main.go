package main

import "fmt"

func main() {
	x := []int{2,3,5,7,11}
	y := x[0:3]
	fmt.Println(y,len(y),cap(y))
	//尾部插入数据
	y = append(y,1,1,1)
	fmt.Println(x,len(x),cap(x))

	c := []int{1,2,3}
	fmt.Println("c",c,len(c),cap(c))
	d := c[:2]
	fmt.Println("d",d,len(d),cap(d))
	var a = []int{1,2,3}
	a = append([]int{0},a...)
	fmt.Println(a)
	//往一个切片头部添加数据都会导致内存重新分配
	a = append([]int{-1,-2,-3},a...)
	fmt.Println(a)

	//中间插入数据
	var aa = []int{-1,-2,3}
	aa = append(aa[:1],append([]int{1,6,6},aa[1:]...)...)
	fmt.Println(aa)

}
