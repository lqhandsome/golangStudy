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

type Queue struct {
	MaxSize int
	arr     [5]*Tree
	head    int
	tail    int
}
//中序和前序非递归方式遍历二叉树
func main() {
	t := resTree()
	//head(t)
	level(t)
}

func head(t *Tree) {
	if t == nil {
		return
	}
	stack := Stack{
		MaxSize: 10,
		Arr:     make([]*Tree,10),
		Top:     -1,
	}

	newT := t
	for newT != nil || stack.Top != -1 {
		//将所有左子节点压入堆栈
		for newT != nil {
			stack.Push(newT)

			newT = newT.Left
		}
		//取出一个节点
		if stack.Top != -1 {
			newT = stack.Pop()
			//查看这个节点的右节点
			newT = newT.Right
		}
	}
}

//层序遍历
func level(t * Tree) {
	if t == nil {
		return
	}
	queue := &Queue{
		MaxSize: 5,
		arr:     [5]*Tree{},
		head:    0,
		tail:    0,
	}
	queue.Push(t)
	for  !queue.IsEmpty() {
		t,_ = queue.Pop()
		fmt.Println(t.Val)
		if t.Left != nil {
			queue.Push(t.Left)
		}
		if t.Right != nil {
			queue.Push(t.Right)
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

//判断队列是否已满
func (queue *Queue) IsFull() bool {
	if (queue.tail + 1) % queue.MaxSize == queue.head {
		return true
	}
	return false
}

//判断队列是否为空
func (queue *Queue) IsEmpty() bool {
	if queue.tail == queue.head {
		return true
	}
	return false
}

//插入一个元素
func (queue *Queue)Push( val *Tree) (err error) {
	if queue.IsFull() {
		fmt.Println("队列已满")
		return errors.New("队列已满")
	}
	queue.arr[queue.tail]  = val
	queue.tail = (queue.tail + 1) % queue.MaxSize
	return nil
}

//获取一个元素
func (queue *Queue)Pop() (val *Tree, er error) {
	if  queue.IsEmpty() {
		fmt.Println("队列为空")
		return nil, errors.New("队列为空")
	}
	val = queue.arr[queue.head]
	queue.head = (queue.head + 1) % queue.MaxSize
	return
}

//判断队列的长度
func (queue *Queue)QueueList() (val int) {
	return (queue.tail + queue.MaxSize - queue.head) % queue.MaxSize
}