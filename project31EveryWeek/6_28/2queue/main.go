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
	que := Queue{
		MaxSize: 5,
	}
	que.Push(1)
	fmt.Println(que.Pop())
	fmt.Println(que)
	fmt.Println(que.arr)
}

//判断队列是否已满
func (queue *Queue) IsFull() bool {
	if  (queue.tail + 1) % queue.MaxSize == queue.head{
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
	queue.arr[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.MaxSize

	return nil
}

//获取一个元素
func (queue *Queue)Pop() (val int, er error) {
	if queue.IsEmpty() {
		fmt.Println("队列为空")
		return -1, er
	}
	val = queue.arr[queue.head]
	queue.head = (queue.head + 1) % queue.MaxSize
	return val, er
}

//判断队列的长度
func (queue *Queue)QueueList() (val int) {
	return (queue.tail + queue.MaxSize - queue.head) % queue.MaxSize
}
