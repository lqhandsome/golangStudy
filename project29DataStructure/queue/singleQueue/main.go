package main

import (
	"errors"
	"fmt"
)

//用数组实现一个非环形队列
//无法复用
type Queue struct {
	maxSize int    //保存最大的数据
	arr     [5]int //保存数据
	front   int    //队列首部
	rear    int    //指向队列尾部
}

func main() {
	queue := Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.PushQueue(val)
			if err != nil {
				fmt.Println("入队列失败", err)
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			val, err := queue.PopQueue()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("取出的数",val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			return
		}

	}
	fmt.Println(queue)
}

func (queue *Queue) PushQueue(val int) (err error) {
	if queue.rear == queue.maxSize-1 {
		return errors.New("队列已满")
	}
	queue.rear++
	queue.arr[queue.rear] = val
	return
}

func (queue *Queue) PopQueue() (val int,err error) {
	if queue.front == queue.rear {
		return -1,errors.New("队列kong")
	}
	queue.front++
	val = queue.arr[queue.front]
	return
}

func (queue *Queue) ShowQueue() {
	fmt.Println("打印队列")
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, queue.arr[i])
	}
	fmt.Println()

}
