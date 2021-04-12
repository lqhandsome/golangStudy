package main

import(
	"fmt"
)

func main()  {
	fmt.Println()

	num := new(int)

	*num = 8
	
	fmt.Printf("%T  %v %v\n",num,num,*num)
}