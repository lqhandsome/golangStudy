package main

import (
	"fmt"
	"strings"
)

func main(){
	//12 返回子串在主串中最后出现的位置
	fmt.Printf("%T %v\n",strings.LastIndex("hello","l"),strings.LastIndex("hello","l"))//int 3

	//13 replace 查找替换字符串strings.Replace("go hello world","go","go语言",n)在哪里找,找什么,替换成什么,替换几次-1全部替换
	fmt.Printf("%T %v\n",strings.Replace("go hello world","go","go语言",-1),strings.Replace("go hello world","go","go语言",-1))

	//14 讲字符串按照指定的字符串拆分为字符串数组
	arr := strings.Split("hello-worl d","-")
	fmt.Printf("%T %v %v\n",arr,arr,len(arr))

	//15 讲字符串转为大写或者小写strings.ToLower ToUpper
		fmt.Println(strings.ToLower("AbCd"),strings.ToUpper("AbCd"))//abcd ABCD

	//16 将字符串中的左右两边空格去除
		fmt.Println(strings.TrimSpace(" abcd "))
	
	//17去除字符串的指定字符串左右两边只要有ab都去除
		fmt.Println(strings.Trim("abcda刘强","ab强"))//cda刘

	//18trimLeft TrimRight
	fmt.Println(strings.TrimLeft("abcda刘强","强"))

	//19判断字符串是否以指定的字符串开头strings.HasPrefix("adv","a")
	fmt.Println(strings.HasPrefix("abcda刘强","a"))

	//20判断字符串是否以定制的字符串结尾 strings.HasSuffix("adv","a")
	fmt.Println(strings.HasSuffix("abcda刘强","强a"))
} 