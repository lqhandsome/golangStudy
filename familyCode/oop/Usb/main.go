package main

import (

	"fmt"
)
type Student struct {

}
func main()  {

	n1 := 1
	n2 := 2.2
	n3 := "string"
	stu1 := Student{}
	stu2 := &Student{}

	TypeJduge(n1,n2,n3,stu1,stu2)
}
func TypeJduge(args ...interface{}){
	for i,v := range args{
		switch v.(type) {
		case bool :
			fmt.Println(i,"bool")
		case string :
			fmt.Println(i,"string")
		case int :
			fmt.Println(i,"int")
		case float64 :
			fmt.Println(i,"float64")
		case Student :
			fmt.Println(i,"Student")
		case *Student :
			fmt.Println(i,"*Student")

		default :
			fmt.Println("default")
		}
	}
}
