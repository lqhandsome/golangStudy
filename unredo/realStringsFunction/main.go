package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main(){

	str := "hello刘强"
	fmt.Println(len(str)) // 11

	fmt.Println([]rune(str)) //[104 101 108 108 111 21016 24378]
	//字符串转为整型
	i1,nil :=strconv.Atoi("123")
	fmt.Printf("%T %v %T\n",i1,i1,nil)

	//整型转为字符串
	fmt.Printf("%T %v\n",strconv.Itoa(123),strconv.Itoa(123))

	//字符串转给[]byte
	fmt.Printf("%T%v\n",[]byte("helloworld"),[]byte("helloworld"))

	//[]byte转为字符串
	ss := string([]byte{97,98,99})
	fmt.Printf("%T %v\n",ss,ss) //string abc

	//10进制转为2,8,16进制
	fmt.Printf("%T %v\n",strconv.FormatInt(123,8),strconv.FormatInt(123,2))//string 1111011

	//查找子串在主串中是否存在
	fmt.Printf("%T %v\n",strings.Contains("hello","o"),strings.Contains("hello","o"))//bool true

	//查找子串在主串中出现的次数
	fmt.Printf("%T %v\n",strings.Count("hello","o"),strings.Count("hello","o"))//int 1

	//不区分大小写比较字符串
	fmt.Printf("%T %v\n",strings.EqualFold("aaa","aAa"),strings.EqualFold("aaa","aAa"))//bool true

	if strings.EqualFold("aaa","aAa") {
		fmt.Printf("%v","吃人")
	}

	//查找子串在主串中第一次出现的位置按照字节
	fmt.Printf("%T %v\n",strings.Index("liuqiang","g"),strings.Index("liuqiang刘强","强"))//int 7
	fmt.Printf("%c\n","liuqiang刘强"[11])

	//查找子串在主串中最后一次出现的位置
	fmt.Println(strings.LastIndex("helloworld","l"))

	//replace 在哪里找,找什么.换成什么,换几次(-1全部替换)
	strRepalce := "hello world"
	fmt.Println(strings.Replace(strRepalce," ","n",-1))

	//将字符串按照指定的字符拆成数组split
	arr := strings.Split("hello-world","-")
	fmt.Printf("%T %v\n",arr,arr)

	//将字串转为全大写或者全小写
	fmt.Println(strings.ToUpper("HelloWorld"),strings.ToLower("HelloWorld"))

	//去除字符串两边的空格
	fmt.Println(strings.TrimSpace(" nihao "))

	//去除字符串左右两边指定字符
	fmt.Println(strings.Trim("helloworld","dh"))

	//去除字符串左或右两边指定字符
	fmt.Println(strings.TrimLeft("helloworld","dh"))
	fmt.Println(strings.TrimRight("helloworld","dh"))

	//查找字符串是否以指定字符开头
	fmt.Println(strings.HasPrefix("a.jpg","a"))

	//查找字符串是否以指定字符开头
	fmt.Println(strings.HasSuffix("a.jpg",".jpg"))
}