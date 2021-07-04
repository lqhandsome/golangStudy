package main

import (
	"errors"
	"fmt"
	"os"
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

type Queue struct {
	MaxSize int
	Arr []*Tree
	Tail int
	Head int
}
//中序和前序非递归方式遍历二叉树
func main() {
	t := resTree()
	//head(t)
	//mid(t)
	level(t)
}

func mid(t *Tree) {
	//初始化一个队列
	stack := Stack{
		MaxSize: 10,
		Arr:     make([]*Tree,10),
		Top:     -1,
	}
	//将头节点压入
	tmp := t
	//stack.Push(tmp)
	for stack.Top != -1 || tmp != nil{
		//把一个节点的所有左子节点压入
		for tmp != nil {
			 stack.Push(tmp)
			 //fmt.Println(tmp.Val)
			 tmp = tmp.Left
		}
		//压入完以后取出一个节点
		newT := stack.Pop()
		if newT != nil {
			fmt.Println(newT.Val)
			if newT.Right != nil {
				tmp = newT.Right
			}
		}
	}
}

func head (t * Tree) {
	//初始化一个队列
	stack := Stack{
		MaxSize: 10,
		Arr:     make([]*Tree,10),
		Top:     -1,
	}
	//将头节点压入
	tmp := t
	//stack.Push(tmp)
	for stack.Top != -1 || tmp != nil{
		//把一个节点的所有左子节点压入
		for tmp != nil {
			stack.Push(tmp)
			fmt.Println(tmp.Val)
			tmp = tmp.Left
		}
		//压入完以后取出一个节点
		newT := stack.Pop()
		if newT != nil {
			if newT.Right != nil {
				tmp = newT.Right
			}
		}
	}
}

//层序遍历
func level(t * Tree) {
	//定义一个队列来保存数据
	queue := Queue{
		MaxSize: 10,
		Arr:     make([]*Tree,10),
		Tail:    0,
		Head:    0,
	}
	queue.Push(t)
	//队列不为空
	for queue.Tail != queue.Head{
		tmp,err := queue.Pop()
		fmt.Println(tmp.Val)
		if err != nil {
			os.Exit(500)
		}
		if tmp.Left != nil {
			queue.Push(tmp.Left)
		}
		if tmp.Right != nil {
			queue.Push(tmp.Right)
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

func (q *Queue) Pop()(val *Tree,err error){
	if q.Tail == q.Head {
		fmt.Println("队列为空")
		return nil, errors.New("队列为空")
	}
	val = q.Arr[q.Head]
	q.Head = (q.Head + 1) % q.MaxSize
	return
}

func (q *Queue) Push(val *Tree) {
	if (q.Tail + 1) % q.MaxSize == q.Head {
		fmt.Println("队列已满")
		return
	}
	q.Arr[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.MaxSize
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
