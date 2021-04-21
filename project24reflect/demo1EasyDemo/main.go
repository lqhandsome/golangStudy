package main

import (
	"fmt"
	"reflect"
)
type stu interface {

}
func main()  {
	var str stu
	str = "1"
	intStr := str.(int)
	fmt.Println(intStr)
	//tt := reflect.TypeOf(str)
	switch str.(type) {
		case bool:
		fmt.Println("bool")
		case string:
		fmt.Println("string")
	}

	fmt.Println()
	fmt.Println(reflect.ValueOf(str))
}