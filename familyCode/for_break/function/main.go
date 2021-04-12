package main

import (
	"fmt"
)

func main(){
	fmt.Println(1)
	fmt.Println(firstFunction(2))
	str := "liuq"
	str = "11"
	fmt.Printf("%T\n",str)
	return 
	for i,v := range str {
		fmt.Println(i,v)
	}
	for i := 0;i < len(str);i++ {
		fmt.Printf("%d = %c\n",i,str[i])
	}
}

func firstFunction( i int) (int ){
	return i + i;
}
