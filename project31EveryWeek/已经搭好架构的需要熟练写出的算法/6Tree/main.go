package main

import (
	"errors"
	"fmt"
)

//用结构体定义树
type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

type Stack struct {
	MaxSize int
	Arr     []*Tree
	Top     int
}

//中序和前序非递归方式遍历二叉树
func main() {
	t := resTree()
	mid(t)
}

func head(t *Tree) {

}

func mid (t * Tree) {

}

//层序遍历
func level(t * Tree) {

}

func (this *Stack) Pop() (Val *Tree) {
	if this.Top == -1 {
		fmt.Println("空栈")
		return nil
	}
	Val = this.Arr[this.Top]
	this.Top--
	return Val
}

func (this *Stack) Push(Val *Tree) error {
	if this.Top == this.MaxSize-1 {
		return errors.New("栈满")
	}
	this.Top++
	this.Arr[this.Top] = Val
	return nil
}

func resTree() *Tree {
	t1 := &Tree{
		Val: 1,
	}
	t2 := &Tree{
		Val: 2,
	}
	t3 := &Tree{
		Val: 3,
	}
	t4 := &Tree{
		Val: 4,
	}
	t5 := &Tree{
		Val: 5,
	}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	return t1
}
