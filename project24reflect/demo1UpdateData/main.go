package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 20
	updateData(&num)
	fmt.Println(num)
}

//修改原值
func updateData(f interface{})  {
	tValue := reflect.ValueOf(f)
	tValue.Elem().SetInt(60)
	fmt.Println(tValue.Kind())
	fmt.Println(tValue.Elem())
	fmt.Printf("%T",tValue.Elem())
}