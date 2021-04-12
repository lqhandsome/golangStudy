package main

import (
	"fmt"
)

func main()  {
	// var input int
	// fmt.Scanln(&input)
	// n := fbnq(input)
	// fmt.Println(n)
	// fmt.Println(monkey(10))
	i := 1
	ii := &i
	c := 'c'
	
	fmt.Printf("%T",c)
	if c == 99 {
		fmt.Println(11)
	}
	fmt.Println(*ii,c)
}

func deth(n int)   {
	if n > 2 {
		n--
		deth(n)
	}
	fmt.Println(n)
}

func detpTwo(n int)  {
	if n > 2 {
		n--
		detpTwo(n)
	} else {
		fmt.Println("n=",n)
	}
}
/*
斐波那契数列 1 1 2 3 5 8
n : 第几个数的值
*/
func fbnq (n int) int {
	if n <= 2 {
		return 1
	}
	return fbnq(n -2) + fbnq(n - 1)
}
/*
f(n)  = f( n -1 ) * 2 + 1
*/
func chengfa (n int) int {
	if n == 1 {
		return 3
	}
	return chengfa( n -1 ) * 2 + 1;
}
/*
猴子吃桃,每天吃昨天剩下的一半并且多吃一个,到n天还剩一个桃子,原来一共有多少个桃子
*/
func monkey(n int) int {
	if n > 1 {
		return ( monkey(n-1) +1) *2
	}
	return 1
}