package main

import (
	"errors"
	"fmt"
)

type node struct {
	val  int
	next *node
}

type Queue struct {
	maxSize int    //保存最大的数据
	arr     [7]int //保存数据
	front   int    //队列首部
	rear    int    //指向队列尾部
}
func main() {

	var prev [...]int
	for  i := 0; i < 8; i++ {
		prev[i] = -1
	}
	fmt.Println(len(prev))
	fmt.Println((prev))
	return
	graph := [...][]int{
		{1,3},
		{0,2,4},
		{1,5},
		{0,4},
		{1,3,5,6},
		{2,4,7},
		{4,7},
		{5,6},
	}



	fmt.Println(graph)
	showList(graph)
}


func bfs(graph [8][]int,s int,t int) {
	var prev [...]int
	for  i := 0; i < 8; i++ {
		prev[i] = -1;
	}
	//保存一个节点是否访问过,访问过为-1
	visited := [len(graph)]int{}

	//用来保存该节点已经被访问,但是指向的节点还没被访问的节点
	queue := &Queue{}
	queue.PushQueue(s)
	visited[s] = -1
}
func showList(graph [8][]int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]) ; j++ {
			fmt.Print(graph[i][j]," ")
		}
		fmt.Println()
	}
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
