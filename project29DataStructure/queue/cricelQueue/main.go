package main

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int //指向队列头
	tail    int //tail指向尾部的下一位
}

//环形队列
func main() {
	queue := CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1、输入add表示添加数据到队列")
		fmt.Println("2、输入get表示弹出队列")
		fmt.Println("3、输入show显示当前队列")
		fmt.Println("4、输入exit退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("入队列失败", err)
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("取出的数", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			return
		}

	}
	fmt.Println(queue)
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.isFull() {
		return errors.New("队列已满")
	}
	this.array[this.tail] = val
	this.tail++
	if this.tail == this.maxSize {
		this.tail = 0
	}
	return
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.isEmpty() {
		return -1, errors.New("队列为空")
	}
	val = this.array[this.head]
	this.head++
	if this.head == this.maxSize {
		this.head = 0
	}
	return
}

//判断环形队列是否为满
func (this *CircleQueue) isFull() (b bool) {
	return (this.tail+1) % this.maxSize == this.head

}

//判断环形队列是否为空
func (this *CircleQueue) isEmpty() (b bool) {
	return this.tail == this.head

}

//环形队列有多个元素
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//显示队列
func (this *CircleQueue) ShowQueue() {
	if this.isEmpty() {
		return
	}
	size := this.Size()
	temp := this.head
	for i := 0;i < size; i++ {
		fmt.Printf("arr[%d]=%d\t",temp,this.array[temp])
		temp = (temp + 1) % this.maxSize
	}
}


