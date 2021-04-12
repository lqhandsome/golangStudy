package main

import (
	"fmt"
	"sort"
)

type Usb interface {
	Say()
}

type Stu struct {

}

func (s *Stu) Say(){
	fmt.Println("*stu")
}
func main() {
	var stu Stu
	var i Usb = &stu

	i.Say()

	var intSlice = []int{2,2542,451,2,15}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
}
