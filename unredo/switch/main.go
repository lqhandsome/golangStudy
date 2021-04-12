package main

import (
	"fmt"
)

func main(){
	var arr =  [...]int{1,2,3,4,5,6}
	//数组指针d
	var pf *[6]int = &arr

	//pf := &arr;fmt.Printf("p =%p t=%T v=%v \n",pf,pf,pf) //p =0xc000012480 t=*[6]int v=&[1 2 3 4 5 6]
	fmt.Printf("p =%p t=%T v=%v \n",pf,pf,pf) //p =0xc000126000 t=*[6]int v=&[1 2 3 4 5 6]`
	fmt.Printf("arr =%p arr=%T arr=%v \n",&arr,arr,arr) //arr =0xc000012480 arr=[6]int arr=[1 2 3 4 5 6]
	return
	// arr := [3]int{1,2,3}
	//s := &arr
	//fmt.Printf("%p %T\n",s,s) //0xc00000c3a0 *[3]int
	//fmt.Printf("%p",&s[0]) //0xc00000c3a0
	//
	//
	//return
	//a := fbnq(6)
	//fmt.Println(a)
	//fmt.Println(N(4))
}
//波非那切
func fbnq(n int)int{
	if n == 1 {
		return 1
	}
	if n ==2 {
		return 1
	}
	return fbnq(n -1) + fbnq(n -2)
}

func N(n int)int {
	if n == 1 {
		return  1
	}
	if n ==2 {
		return  2
	}
	return n * N(n-1)
}