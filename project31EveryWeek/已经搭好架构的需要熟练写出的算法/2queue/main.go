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

}

//判断队列是否已满
func (queue *Queue) IsFull() bool {
}

//判断队列是否为空
func (queue *Queue) IsEmpty() bool {
}

//插入一个元素
func (queue *Queue)Push( val int) (err error) {
}

//获取一个元素
func (queue *Queue)Pop() (val int, er error) {
}

//判断队列的长度
func (queue *Queue)QueueList() (val int) {
}
