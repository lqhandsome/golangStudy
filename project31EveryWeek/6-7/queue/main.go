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
	queue := &Queue{
		MaxSize: 5,
		arr:     [5]int{},
		head:    0,
		tail:    0,
	}

	fmt.Println(Push(queue, 1))
	fmt.Println(Push(queue, 2))
	fmt.Println(Push(queue, 3))
	fmt.Println(Push(queue, 4))
	fmt.Println(Push(queue, 5))
	fmt.Println(QueueList(queue))
	fmt.Println(queue)
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(Pop(queue))
	fmt.Println(queue)

}

//判断队列是否已满
func IsFull(queue *Queue) bool {
	//队列尾部的下一位就是头节点，代表队列满了
	if (queue.tail+1)%queue.MaxSize == queue.head {
		return true
	} else {
		return false
	}
}

//判断队列是否为空
func IsEmpty(queue *Queue) bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

//插入一个元素
func Push(queue *Queue, val int) (err error) {
	if IsFull(queue) {
		return errors.New("队列已满")
	}
	queue.arr[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.MaxSize
	return nil
}

//获取一个元素
func Pop(queue *Queue) (val int, er error) {
	if IsEmpty(queue) {
		return -1, errors.New("队列为空")
	}
	val = queue.arr[queue.head]
	queue.arr[queue.head] = 0
	queue.head = (queue.head + 1) % queue.MaxSize
	return val, er
}

//判断队列的长度
func QueueList(queue *Queue) (val int,err error) {
	if IsEmpty(queue) {
		return 0, err
	}
	val = (queue.tail + queue.MaxSize - queue.head) % queue.MaxSize
	return val, err
}
