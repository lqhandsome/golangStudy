package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v float64 = 1.2
	transfor(v)
}


func transfor(this interface{}) {
	rValue := reflect.ValueOf(this)
	rType := reflect.TypeOf(this)
	fmt.Printf("type %v,kind=%v，值%v",rType,rValue.Kind(),rValue.Float())
	res := rValue.Interface()
	resres := res.(float64)
	fmt.Println(resres)
}
