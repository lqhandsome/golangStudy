package main

import (
	"fmt"
)

func main()  {
	// fmt.Println(11)
	n1 := 2
	n2 := 1
	swap(&n1,&n2)
	fmt.Println(n1,n2)
	a := 5
	b := 7
	fmt.Println(testYinyong(&a,&b))
}

func swap(n1 *int,n2 *int){
		*n1 = *n2 + *n1
		*n2 = *n1 - *n2
		*n1 = *n1 - *n2
}
func testYinyong(n1 *int,n2 * int) (int,int,*int){
	a := *n1
	b := *n2
	return a ,b,n1
}