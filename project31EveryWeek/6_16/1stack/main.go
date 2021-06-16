package main

import "fmt"

type Stack struct {
	arr []int
	top int
}
func main() {
	stack := Stack{
		make([]int,5),
		-1,
	}
	stack.push(1)
	fmt.Println(stack.pop())
}
func (this *Stack) pop() int {
	if this.top == -1 {
		return  -1
	}
	val := this.arr[this.top]
	this.top--
	return val
}
func (this *Stack) push(val int)  {
	this.top++
	this.arr[this.top] = val
}
