package main

import (
	"fmt"
)

func main(){
	res := myFun(getSum,10,20)
	fmt.Println(res)
	funcOne  := getSum
	fmt.Println(funcOne(20,20))
}

func getSum(n1 int,n2 int) int {
	return n1 + n2
}

func myFun(funcvat func(int,int) int ,n1 int,n2 int) int {
	return funcvat(n1 ,n2)
}