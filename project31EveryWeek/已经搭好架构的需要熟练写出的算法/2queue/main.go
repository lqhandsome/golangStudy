package main

import (
	"errors"
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

}

//判断队列是否已满
func IsFull(queue *Queue) bool {

}

//判断队列是否为空
func IsEmpty(queue *Queue) bool {

}

//插入一个元素
func Push(queue *Queue, val int) (err error) {

}

//获取一个元素
func Pop(queue *Queue) (val int, er error) {

}

//判断队列的长度
func QueueList(queue *Queue) (val int,err error) {

}
