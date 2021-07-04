package main

import "fmt"

type Stack struct {
	arr []int
	top int
}
func main() {
	stack := &Stack{
		arr: make([]int,10),
		top: -1,
	}
	stack.push(11)
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack)
}
func (this *Stack) pop() int {
	if this.top == -1 {
		fmt.Println("没有元素")
		return -1
	}
	res := this.arr[this.top]
	this.top--
	return res
}
func (this *Stack) push(val int)  {
	this.top++
	this.arr[this.top] = val
	return
}
