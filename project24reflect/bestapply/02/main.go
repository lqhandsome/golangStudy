package main

import (
	"fmt"
	"reflect"
)
type Monster struct {
	Name string
	Age int
}

func main() {
	monster := Monster{
		"牛魔王",
		500,
	}
	testStruct(&monster)
	fmt.Println(monster)
}

//遍历结构体并修改值
func testStruct(c interface{}) {

	//typ := reflect.TypeOf(c)

	val := reflect.ValueOf(c)

	if val.Kind() != reflect.Ptr {
		fmt.Println("类型错误")
		return
	}
	varRes := val.Interface()
	res := varRes.(*Monster)
	fmt.Println(varRes)
	fmt.Println(res.Name)
	res.Name = "好孩子"
}
