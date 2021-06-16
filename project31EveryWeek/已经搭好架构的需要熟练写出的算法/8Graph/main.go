package main

import (
	"errors"
)

type Graph struct {
	val  int
	node []int
}

type Queue struct {
	maxSize int    //保存最大的数据
	arr     [8]int //保存数据
	front   int    //队列首部
	rear    int    //指向队列尾部
}

func main() {
	graph := [...][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5, 6},
		{2, 4, 7},
		{4, 7},
		{5, 6},
	}
	//bfs(graph, 0, 6) //[-1 0 1 0 1 2 4 -1]

	dfs(graph, 0, 6) //-1 0 1 4 5 2 7 5   [-1 0 1 4 5 2 4 5]

}

func bfs(graph [8][]int, t int, s int) {
}

var found bool

func dfs(graph [8][]int, t int, s int) {

}
func singDfs(graph [8][]int, w int, s int, visited *[8]bool, prev []int) {


}

func (queue *Queue) PushQueue(val int) (err error) {
	if queue.rear == queue.maxSize-1 {
		return errors.New("队列已满")
	}
	queue.arr[queue.rear] = val
	queue.rear++
	return
}

func (queue *Queue) PopQueue() (val int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("队列kong")
	}
	val = queue.arr[queue.front]
	queue.front++
	return
}
