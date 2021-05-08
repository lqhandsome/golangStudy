package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	array [5]int
	head int
	tail int
}


func main() {
	queue := Queue{
		MaxSize: 5,
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
}

//判断队列是否满了
func (this *Queue) IsFull() bool {
	return (this.tail + 1 + this.MaxSize ) % this.MaxSize == this.head
}

//判断队列是否为空
func (this *Queue) IsEmpty() bool {
	return this.tail == this.head
}

//判断队列的个数

func (this *Queue) QueueSize() int {
	return (this.tail + this.MaxSize - this.head) % this.MaxSize
}

func (this *Queue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.MaxSize
	return
}

func (this *Queue) Pop() (val int,err error) {
	if this.IsEmpty() {
		return -1,errors.New("队列为空")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.MaxSize
	return
}

func (this *Queue) ShowQueue() {
	if this.IsEmpty() {
		fmt.Println("队列为空")
		return
	}
	tmp := this.head
	size := this.QueueSize()
	for i := 0 ; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t",tmp,this.array[tmp])
		tmp = (tmp + 1 ) % this.MaxSize
	}
	fmt.Println()

}