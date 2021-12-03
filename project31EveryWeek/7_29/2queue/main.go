package main

import (
	"errors"
	"fmt"
	"log"
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
	queue.IsFull()
	queue.Push(1)
	fmt.Println(queue.Pop())

}

//判断队列是否已满
func (queue *Queue) IsFull() bool {

	if (queue.tail + 1 + queue.MaxSize) % queue.MaxSize == queue.head {
		fmt.Println("队列已满",queue.tail,queue.head)
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
		fmt.Println("队列已满，无法加入")
		return errors.New("队列满")
	}
	queue.arr[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.MaxSize
	return nil
}

//获取一个元素
func (queue *Queue)Pop() (val int, er error) {
	if queue.IsEmpty() {
		log.Fatalln("队列为空无法弹出")
		return -1, errors.New("队列空")
	}
	val = queue.arr[queue.head]
	queue.head = (queue.head + 1) % queue.MaxSize
	return
}

//判断队列的长度
func (queue *Queue)QueueList() (val int) {
	return (queue.tail + queue.MaxSize - queue.head) % queue.MaxSize
}
