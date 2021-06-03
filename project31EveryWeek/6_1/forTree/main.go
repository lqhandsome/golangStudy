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
	q := Stack{
		MaxSize: 50,
		Top:     -1,
		Arr:     make([]*Tree, 10),
	}
	if t == nil {
		fmt.Println("空tree")
		return
	}
	tmpT := t
	for tmpT != nil || q.Top != -1 {
		for tmpT != nil {
			q.Push(tmpT)
			fmt.Println(tmpT.Val)
			tmpT = tmpT.Left
		}

		if q.Top != -1 {
			tmpT = q.Pop()
			tmpT = tmpT.Right
		}

	}
}

func mid (t * Tree) {
	if t == nil {
		fmt.Println("空树")
		return
	}

	s := Stack{
		MaxSize: 50,
		Arr: make([]*Tree,50),
		Top: -1,
	}
	tmpT := t
	for tmpT != nil || s.Top != -1 {
		//把左子树全部压入
		for tmpT != nil {
			s.Push(tmpT)

			tmpT = tmpT.Left
		}
		//压入完成之后取出一个，就是最后最后左子树
		if s.Top != -1 {
			tmpT = s.Pop()
			//遍历右子树之前输出节点的值
			fmt.Println(tmpT.Val)
			//遍历右子树
			tmpT = tmpT.Right
		}
	}
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
