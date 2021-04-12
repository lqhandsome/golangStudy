package main
import (
	"fmt"
	"unsafe"
)

var (
	v3 = "4"
	v4 = "5"
)
//基本数据类型
	/*
	数值型 
		整型 int,int8(8位 范围 -127 ~ 126) uint8(无符号),byte(一个字节)
		浮点型 float32(32位),float64
	字符型(ascii) a,b,c,d 单个字符
	布尔型 true or false
	字符串
	*/
//派生型
	/*
	指针
	数组
	结构体 类似于class
	管道
	函数
	切片
	接口
	map
	*/

	/*
		int8 一个字节八位,表示范围 -128~127(2^7)一位表示正负 为什么正数到127因为-0不表示所以多一个
		uint8 占用一个字节八位,表示范围0~255(2^8 -1)为什么减一 因为0要占用一个
		int,uint,rune,byte
		int 32位系统4个字节 -2^31 ~ 2^31 - 1 64位系统-2^63 ~ 2^63 - 1
		uint 32位系统4个字节 0 ~ 2^32 -1 64位系统0 ~ 2^64 -1
		rune 与int32一样 -2^31 ~ 2^31 - 1 等价int32 用来表示unicode
		byte 与 uint8等价 0~255 当要储存字符时选用byte
	*/
func main()  {
	i := 1
	var ii int8
	ii = 127
	var v1 uint
	v1 = 18446744073609551611
	n1 := 100
	fmt.Println("v3,v4",i,ii,'1',v1)
	fmt.Printf("n1的类型%T",n1)
	var n2 int = 10
	fmt.Printf("n2的类型是 %T n2占用的字节数是 %d",n2,unsafe.Sizeof(n2))
}