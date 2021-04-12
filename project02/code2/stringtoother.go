package main

import	(
	"fmt"
	"strconv"
)
//strconv的Itoa讲int转为string
// func main()  {
// 	var i int64 = 100
// 	var str string
// 	str = strconv.FormatInt(int64(i),10)
// 	fmt.Println(str)
// }

//string转为其他数据类型
func main()  {
	var str string = "111"
	var b bool
	var i int64
	var f float64
	b,_ = strconv.ParseBool(str)
	fmt.Printf("%T%v\r",b,b)

	i,_ = strconv.ParseInt(str,10,0)
	
	fmt.Printf("%T%v\n",i,i)

	T := int(i)
	fmt.Printf("%T%v\n",T,T)
	
	f,_ = strconv.ParseFloat("3.14",64)
	fmt.Printf("%T%v\n",f,f)

	e := float32(f)
	fmt.Printf("%T%v\n",e,e)
	
}
