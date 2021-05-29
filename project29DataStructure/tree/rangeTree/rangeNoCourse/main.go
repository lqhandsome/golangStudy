package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxSize int
	Arr     []*Tree
	Top     int
}

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

type Queue struct {
	MaxSize int
	arr     [10]*Tree
	head    int
	tail    int
}

func main() {
	t := resTree()
	//midTree(t) //42513
	//中序遍历
	//rangeCourseTreeAInner(t)
	//前序遍历
	//rangeCourseTreeFront(t)
	//后续遍历
	fmt.Println("-----------------------------")
	//rangeCourseTreeAfter(t)
	//层序遍历rabbitmqcli-560-rabbit@60a2b7bdafe5
	rangeCourseTreeLeve(t)
}

/**
前序遍历，主要通过栈保存，该节点的子节点，当遍历到底部，输出保存的节点docker exec -it rabbitmqCluster01 bash
*/
func rangeCourseTreeFront(t *Tree) {
	s := Stack{
		MaxSize: 50,
		Top:     -1,
		Arr:     make([]*Tree, 10),
	}
	tmpT := t
	for tmpT != nil || s.Top != -1 {
		for tmpT != nil {
			s.Push(tmpT)
			fmt.Println(tmpT.Val)
			tmpT = tmpT.Left
		}
		if s.Top != -1 {
			tmpT = s.Pop()
			tmpT = tmpT.Right
		}
	}
}

//层序遍历
func rangeCourseTreeLeve(t *Tree) {
	if t == nil {
		fmt.Println("空树")
		return
	}
	queue := &Queue{
		MaxSize: 10,
	}
	queue.Push(t)
	for queue.tail != queue.head {
		tmp,err := queue.Pop()
		returnErr(err,"74")
		fmt.Println(tmp.Val)
		if tmp.Left != nil {
			queue.Push(tmp.Left)
		}
		if tmp.Right != nil {
			queue.Push(tmp.Right)
		}
	}
}
func rangeCourseTreeAfter(t *Tree) {
	s := Stack{
		MaxSize: 50,
		Top:     -1,
		Arr:     make([]*Tree, 10),
	}
	p := &Tree{}
	for t != nil || s.Top != -1 {
		for t != nil {
			s.Push(t.Left)
			t = t.Left
		}
		if s.Top != -1 {
			t = s.Pop()
			fmt.Println(t)
			if t == nil || t == p {
				p = t
				t = nil
			} else {
				fmt.Println(72)
				s.Push(t)
				t = t.Right
			}
		}
	}
}

func rangeCourseTreeAInner(t *Tree) {
	s := Stack{
		MaxSize: 50,
		Top:     -1,
		Arr:     make([]*Tree, 10),
	}
	tmpT := t
	for tmpT != nil || s.Top != -1 {
		for tmpT != nil {
			s.Push(tmpT)
			tmpT = tmpT.Left
		}
		if s.Top != -1 {
			tmpT = s.Pop()
			fmt.Println(tmpT.Val)
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
func midTree(t *Tree) {
	if t != nil {
		midTree(t.Left)
		fmt.Println(t.Val)
		midTree(t.Right)
	}
}


//队列插入
func (q *Queue) Push(val *Tree) (err error) {
	if (q.tail+1)%q.MaxSize == q.head {
		return errors.New("队列已满")
	}
	q.arr[q.tail] = val
	q.tail = (q.tail + 1) % q.MaxSize
	return
}

//队列弹出
func (q *Queue) Pop() (val *Tree, err error) {
	if q.head == q.tail {
		return nil, errors.New("队列是空")
	}
	val = q.arr[q.head]
	q.head = (q.head + 1) % q.MaxSize
	return
}
func returnErr(err error, message string) {
	if err != nil {
		fmt.Println(err, message)
	}
}