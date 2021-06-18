package main

import (
	"errors"
	"fmt"
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
	if t == s {
		return
	}
	//定义queue用来保存还没被访问的节点
	queue := Queue{
		maxSize: 8,
		arr:     [8]int{},
		front:   0,
		rear:    0,
	}
	//判断一个节点是否被访问过
	visited := [8]bool{}
	//prev 用来保存路径
	prev := [8]int{-1,-1,-1,-1,-1,-1,-1,-1}

	//将头节点压入
	queue.PushQueue(t)
	visited[t] = true
	for queue.front != queue.rear {
		//取出一个顶点
		w,err := queue.PopQueue()
		if err != nil {
			fmt.Println("队列错误")
			return
		}
		visited[w] = true
		//遍历这个点的所有顶点
		for i :=0;i < len(graph[w]) && w != -1 ;i++{
			//取出一条边
			tmp := graph[w][i]
			//这个节点没有被访问过
			if !visited[tmp] {
				prev[tmp] = w
				//这个节点从哪里来的
				if tmp == s {
					fmt.Println(prev)
					return
				}
				err = queue.PushQueue(tmp)
				if err != nil {
					fmt.Println("压入队列出错",err,queue.arr)
					//return
				}
				visited[tmp] = true
			}
		}
	}
}

var found bool

func dfs(graph [8][]int, t int, s int) {
	if t == s {
		return
	}

	//用来记录这个节点是否访问过
	visited := [8]bool{}
	//用来保存路径
	prev := make([]int,8)
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}
	visited[t]  = true
	singDfs(graph,t,s,&visited,prev)
	fmt.Println(prev)
}
func singDfs(graph [8][]int, w int, s int, visited *[8]bool, prev []int) {
	if found {
		return
	}
	for i := 0 ; i < len(graph[w]);i++ {
		tmp := graph[w][i]
		if !visited[tmp] {
			prev[tmp] = w
			if tmp == s {
				found = true
				return
			}
			visited[tmp] = true
			singDfs(graph,tmp,s,visited,prev)
		}
	}
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
