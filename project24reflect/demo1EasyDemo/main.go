package main

import (
	"fmt"
	"reflect"
)

func main()  {
	b := 100
	reflectBaseData(b)

	stu := Student{
		"lq",
		23,
	}
	reflectStruct(stu)
}

type Student struct {
	Name string
	Age  int

}



//结构体的反射
func reflectStruct(r interface{}) {
	//通过反射获取的传入的变量的type，kind，值
	//获取到reflect.TypeOf()
	rTyp := reflect.TypeOf(r)
	fmt.Println("rTyp=",rTyp)


	//2.获取到reflect.Value
	rValue := reflect.ValueOf(r)
	fmt.Println("rValue=",rValue)

	//下面将rVal转成interface{} 但是并不能访问到字段
	iV := rValue.Interface()
	fmt.Printf("iV=%v iV type=%T \n",iV,iV)
	iVType := fmt.Sprintf("%T\n",iV)
	fmt.Printf("iVType type = %T",iVType)
	//需要使用类型断言才能访问到
	a ,ok:= iV.(Student)
	fmt.Printf("a=%v a type=%T name=%v ok=%v\n",a,a,a.Name,ok)
	fmt.Printf("rTyp kind=%v",rTyp.Kind())

}

//基本数据类型 原始基本数据->reflect->interface{}->原始数据
func reflectBaseData(b interface{}) {
	//查看b的类型,打印并不是真正的类型
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)


	//2.获取到reflect.Value，打印出来的并不是真正的值
	rValue := reflect.ValueOf(b)
	fmt.Println(rValue)
	num1 := 20
	//需要转换才能使用，万一需要使用到其他的数据类型呢？这就是使用到类型断言了
	fmt.Println(int64(num1) + rValue.Int())
	//类型断言
	//1.先将数据转为interface{}类型
	iV := rValue.Interface()
	//2.将interface{}转为需要的类型
	res := iV.(int)
	var num2 int = 100
	fmt.Println(res + num2)
}

