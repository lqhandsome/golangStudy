package main

import (
	"fmt"
	"reflect"
)

//定义一个结构体
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32
	Sex string
}
func main() {
 monster := Monster{
 	Name: "红孩儿",
 	Age: 356,
 	Score: 566.3,
 	Sex: "man",
 }
	testStruct(monster)
}

func (this Monster) One(){
	fmt.Println("我是第一个方法")
	this.Name = "牛魔王"
}

func (this Monster) Two(a interface{},b interface{}){
	fmt.Println(a.(int) + b.(int))
}
func testStruct(this interface{}){
	rType := reflect.TypeOf(this)

	rValue := reflect.ValueOf(this)

	kd := rType.Kind()
	//判断结构体是不是struct类型的
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//通过反射获取结构体有多少个字图案
	fieldNums := rValue.NumField()
	fmt.Println("struct 一共有",fieldNums,"个字段")

	for i := 0; i < fieldNums; i++ {
			fmt.Printf("field %d: 值为%v\n",i,rValue.Field(i))
			tagVal := rType.Field(i).Tag.Get("json")
			if tagVal != ""{
				 fmt.Printf("Field %d:tag为%v\n",i+1,tagVal)
			}
	}

	//获取结构体有多少个方法
	methodNums := rValue.NumMethod()
	fmt.Println(methodNums)

	//方法的排序默认是按照函数名的排序（ASCII）
	rValue.Method(0).Call(nil)

	//调用传参参数格式[]reflect.Value
	one := reflect.ValueOf(10)
	two := reflect.ValueOf(20)
	args := []reflect.Value{one,two}
	rValue.Method(1).Call(args)
}
