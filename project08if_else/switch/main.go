package main

import (
	"fmt"
)

func main()  {
	// var i int = 10
	// switch iii := 10; {
	// 	case iii == 10:
	// 		fmt.Println("10")
	// 		//穿透
	// 		fallthrough
	// 	case i > 20:
	// 		fmt.Println(20)
	// }

	// if ii :=10; ii<=10{
	// 	fmt.Println(111)
	// }	

	// fmt.Printf("%d%v%T%c",a,a,a,a) --- 97 97 uint8 a
	//练习一把小写的char型转化为大写(键盘输入)
	// var c byte 
	
	// fmt.Println("请输入一个字符")
	// fmt.Scanf("%c",&c)
	/* 方法一
		
	var a int
	
	fmt.Println("请输入一个字符")
	fmt.Scanln(&a)
	switch c {
		case 'a':
			c = 'A'
			fmt.Printf("%c",c)
		case 'b':
			c = 'B'
			fmt.Printf("%c",c)
		case 'c':
			c = 'C'
			fmt.Printf("%c",c)
		case 'd':
			c = 'D'
			fmt.Printf("%c",c)
		case 'e':
			c = 'E'
			fmt.Printf("%c",c)
		default:
			fmt.Println("other")
	}*/
	//方法二
	// switch {
	// 	case c >= 97 && c <=101 :
	// 		 c -= 32
	// 		 fmt.Printf("%c",c)
	// 	default:
	// 		fmt.Println("其他")
	// }


	var month byte
	fmt.Println("请输入一个月份")
	fmt.Scanln(&month)
	switch month {
	case 3,4,5,97:
		fmt.Println("春季")
		fmt.Printf("%c",month)
	case 6,7,8:
		fmt.Println("夏季")
	case 9,10,11:
		fmt.Println("秋季")
	case 12,1,2:
		fmt.Println("冬季季")
	}

	
	
}