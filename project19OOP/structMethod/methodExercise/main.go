package main

import "fmt"


type MethodUtils struct {

}
func main() {
	m := MethodUtils{}
	m.print(5,6)
}

func (M MethodUtils) print(m int,n int){
	for i := 0;i < m;i++ {
		for j := 0;j < n ;j++ {
			fmt.Print(1," ")
		}
		fmt.Println()
	}
}