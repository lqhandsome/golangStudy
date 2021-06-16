package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world刘1"
	fmt.Println(len(str))
	fmt.Println(str[1:])
	for i:=0;i <len(str);i++{
		fmt.Printf("%c\n",str[i])
	}

	for _,val:=range str {
		fmt.Printf("%c\n",val)
	}
	run := []rune(str)
	var a string
	for _,val:=range run {
		a += string(val)
		fmt.Printf("%c\n",val)
	}
	fmt.Println(a)
	//讲字符串转为整型
	strInt := "1000"
	fmt.Println(strconv.Atoi(strInt))
	//检查一个字符串在另个字符串那种是否存在
	fmt.Println(strings.Contains("helllo","lo"))
	//查找一个字符串在另个字符串中出现的次数
	fmt.Println(strings.Count("hello","l"))
	//比较两个字符串不区分大小写
	fmt.Println(strings.EqualFold("abc","ABc"))
	//返回子串在主串中出现的位置，没有则返回-1
	fmt.Println(strings.Index("abcda","a"))
	//返回子串在主串中最后出现的位置，没有则返回-1
	fmt.Println(strings.LastIndex("abcda","a"))
	//重复字符串
	fmt.Println(strings.Repeat("a",4))
	//查找替换字符串
	fmt.Println(strings.Replace("abcd","bc","ee",2))
	//查找替换字符串所有
	fmt.Println(strings.ReplaceAll("abbbcd","b","ee"))
	fmt.Println(strings.Replace("abbbcd","b","ee",strings.Count("abbbcd","b")))
	//讲字符串拆分为数组
	fmt.Println(strings.Split("a-b-c-d","-"))
	//讲字符串转为小写
	fmt.Println(strings.ToLower("ABCs"))
	//讲字符串转为大写
	fmt.Println(strings.ToUpper("ABCs"))
	//去除字符串两边空格
	fmt.Println(strings.Trim(" abcd "," "))
	fmt.Println(strings.TrimSpace(" abcd "))
}
