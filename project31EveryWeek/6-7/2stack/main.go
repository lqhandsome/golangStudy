package main

import "fmt"

type Stack struct {
	arr []int
	top int
}

func main() {
	s := Stack{
		arr: []int{-1,-1,-1,-1,-1,-1,-1},
		top: -1,
	}
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.arr)
	fmt.Println(s.pop())
}

func (this *Stack) pop() int {
	if  this.top == -1{
		return -1
	}
	val := this.arr[this.top]
	this.top--
	return  val
}
func (this *Stack) push(val int)  {
	this.top++
	this.arr[this.top] = val
	return
}
