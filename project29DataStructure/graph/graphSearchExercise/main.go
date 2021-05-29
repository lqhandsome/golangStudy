package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	arr     [10]int
	head    int
	tail    int
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
	//bfs(graph,3,6)
	//BFS(graph, 0, 6)
	DFS(graph, 0, 6)
}

var found bool

func DFS(graph [8][]int, t int, s int) {
	if s == t {
		return
	}
	//用来保存路径
	prev := make([]int, 8)
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}
	//用来保存节点是否访问过
	visited := [8]bool{}
	visited[t] = true

	SDFS(graph, t, s, prev, &visited)
	fmt.Println(prev) //
}
func SDFS(graph [8][]int, t int, s int, prev []int, visited *[8]bool) {
	visited[t] = true
	if t == s {
		return
	}
	for i := 0; i < len(graph[t]); i++ {
		if t == 7 {
			fmt.Println(57,graph[t][i],visited[6])
		}
		tmp := graph[t][i]
		if !visited[tmp] {
			prev[tmp] = t
			SDFS(graph,tmp,s,prev,visited)
		}
	}
}

func BFS(graph [8][]int, t int, s int) {
	if t == s {
		return
	}

	//定义prev用来保存路径
	prev := [8]int{-1, -1, -1, -1, -1, -1, -1, -1}

	//定义visited保存节点是否被访问过
	visited := [8]bool{}

	//定义queue用来保存改节点被访问但是子节点没有被访问的节点
	queue := Queue{
		MaxSize: 10,
		arr:     [10]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		head:    0,
		tail:    0,
	}
	queue.Push(t)
	visited[t] = true
	for queue.tail != queue.head {
		w, err := queue.Pop()
		returnErr(err, "51")
		for i := 0; i < len(graph[w]); i++ {
			tmp := graph[w][i]
			if !visited[tmp] {
				prev[tmp] = w
				if tmp == s {
					fmt.Println(prev)
					return
				} else {
					visited[tmp] = true
					err = queue.Push(tmp)
					returnErr(err, "65")
				}
			}
		}
	}
}

//队列插入
func (q *Queue) Push(val int) (err error) {
	if (q.tail+1)%q.MaxSize == q.head {
		return errors.New("队列已满")
	}
	q.arr[q.tail] = val
	q.tail = (q.tail + 1) % q.MaxSize
	return
}

//队列弹出
func (q *Queue) Pop() (val int, err error) {
	if q.head == q.tail {
		return -1, errors.New("队列是空")
	}
	val = q.arr[q.head]
	q.head = (q.head + 1) % q.MaxSize
	return
}
func returnErr(err error, message string) {
	if err != nil {
		fmt.Println(err, message)
	}
}


