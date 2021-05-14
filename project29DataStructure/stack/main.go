package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    //最大存放个数
	Top    int    //表示栈顶
	arr    [5]int //存放数据
}

//入栈
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop - 1 {
		fmt.Println("栈满")
		return errors.New("栈满")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈为空")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Println(this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (err error, val int) {
	if this.Top == -1 {
		return errors.New("栈为空"), -1
	}
	val = this.arr[this.Top]
	this.Top--
	return
}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, //没有数据
	}
	stack.Push(1)
	stack.Push(5)
	stack.List()
	fmt.Println(stack.Pop())
	stack.List()
}
