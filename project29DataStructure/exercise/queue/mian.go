package main

import "fmt"

type Queue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

func main() {
	queue := &Queue{
		maxSize: 5,
	}
	Push(queue,3)
	Push(queue,2)
	Push(queue,1)
	fmt.Println(queue)
	//fmt.Println(QueueSize(queue))
	ShowQueue(queue)
	Pop(queue)
	Pop(queue)
	fmt.Println()
	ShowQueue(queue)
}

//判断队列是否为空
func IsEmpty(queue *Queue) bool {
	return queue.head == queue.tail
}

//判断是否满
func IsFull(queue *Queue) bool {
	return (queue.tail+1)%queue.maxSize == queue.head
}

//队列的长度
func QueueSize(queue *Queue) int {
	if IsEmpty(queue) == true {
		return 0
	}
	return (queue.tail + queue.maxSize - queue.head) % queue.maxSize
}

func ShowQueue(queue *Queue) {
	queueSize := QueueSize(queue)

	tmp := queue.head
	for i := 0; i < queueSize; i++ {
		fmt.Println(queue.array[tmp])
		tmp = (tmp + 1) % queue.maxSize
	}
}

func Push(queue *Queue,val int) bool {
	if IsFull(queue) {
		return false
	}
	queue.array[queue.tail] = val
	queue.tail = (queue.tail + 1) % queue.maxSize
	return true
}

func Pop(queue *Queue) int  {
	if IsEmpty(queue) {
		return  -1
	}

	val := queue.array[queue.head]
	queue.head = (queue.head + 1) % queue.maxSize
	return val
}