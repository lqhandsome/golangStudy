package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	arr     [5]int
	head    int
	tail    int
}

/**
实现一个环形队列使用数组
*/
func main() {
	queue := Queue{
		MaxSize: 5,
		arr:     [5]int{},
		head:    0,
		tail:    0,
	}
	queue.Push(100)
	queue.Push(100)
	queue.Push(100)
	queue.Push(100)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	queue.Push(100)
	fmt.Println(queue.QueueList())
	fmt.Println(queue.IsFull())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.tail)
	fmt.Println(queue.head)
	fmt.Println(queue.QueueList())
	fmt.Println(queue.Pop())
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
func (queue *Queue)Push( val int) (err error) {
	if queue.IsFull() {
		fmt.Println("队列已满")
		return errors.New("队列已满")
	}
	queue.arr[queue.tail]  = val
	queue.tail = (queue.tail + 1) % queue.MaxSize
	return nil
}

//获取一个元素
func (queue *Queue)Pop() (val int, er error) {
	if  queue.IsEmpty() {
		fmt.Println("队列为空")
		return -1, errors.New("队列为空")
	}
	val = queue.arr[queue.head]
	queue.head = (queue.head + 1) % queue.MaxSize
	return
}

//判断队列的长度
func (queue *Queue)QueueList() (val int) {
	return (queue.tail + queue.MaxSize - queue.head) % queue.MaxSize
}
