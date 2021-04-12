package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	//1.    len()返回变量的字节长度 str[1] 得到是ascii值
	//一个ascill一个字节,中文三个字节
	fmt.Println(len("aaa刘"))
	str := "abcd刘强"
	//2.	rune 可以讲字符串按照字符为一个长度的
	r := []rune(str)

	fmt.Println(len(r))  // 6
	fmt.Printf("%T%v\n",r,r) //[]int32[97 98 99 100 21016 24378]

	//3.	strconv
	//3.1 Atoi 字符串转为整形  返回int,nii 报错也会转为0
	i,_ :=strconv.Atoi("111a")
	fmt.Println(i)

	//3.2 Itoa 整形转为字符串 返回string 一定会转化成功
	str3 := strconv.Itoa(1234)
	fmt.Printf("%T%v\n",str3,str3) //string1234

	//4 字符串转为 []byte var bytes = []byte("hello go")
	fmt.Printf("%T%v\n",[]byte("hello go"),[]byte("hello go"))//[]uint8[104 101 108 108 111 32 103 111]

	//5 []byte转为字符串 str := string([]byte{97,98,99})
	str5 := string([]byte{97,98,99})
	fmt.Printf("%T%v\n",str5,str5)

	//6 10进制转为2,16,8进制 str := strconv.FormatInt(123,2),返回对应的字符串
	str6 := strconv.FormatInt(123,16)
	fmt.Printf("%T%v\n",str6,str6)

	//7 查找子串是否在指定的字符串中strings.Contains("lqhandsome","lq") true
	fmt.Printf("%v\n",strings.Contains("lqhandsome","lq"))

	//8 统计一个串在另一个串中的个数 strings.Count("lqhello","l") int 3
	i8 := strings.Count("lqhello","l")
	fmt.Printf("%T%d\n",i8,i8)

	//9 不区分大小写的字符串比较(==是区分大小写的):fmt.Printf(strings.EqualFold("abc","ABC"))
	fmt.Printf("%T%t\n",strings.EqualFold("abc","ABC"),strings.EqualFold("abc","ABC"))
	
	//10 返回字符串在某字符串的位置 strings.Index("lqhandsome","q")
	fmt.Println(strings.Index("lqhandsome刘强","刘")) //10 没有则返回-1
}

