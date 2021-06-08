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
	//定义visited数组用来保存元素是否访问过
	visited := [8]bool{}

	//定义一个数组用来记录节点从哪里来的
	prev := [8]int{}
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}
	visited[t] = true
	//定义队列用来保存节点访问过但是没有被访问过的子节点
	queue := Queue{
		maxSize: 8,
		arr:     [8]int{-1, -1, -1, -1, -1, -1, -1},
		front:   0,
		rear:    0,
	}
	//将要查找的先压入
	queue.PushQueue(t)
	for queue.front != queue.rear {
		//取出一个
		w, err := queue.PopQueue()
		if err != nil {
			fmt.Println("队列弹出错误")
			return
		}
		//查看这个边
		for i := 0; i < len(graph[w]) && w != -1; i++ {
			tmp := graph[w][i]
			if !visited[tmp] {
				prev[tmp] = w
				if tmp == s {
					fmt.Println(prev)
					return
				}
				visited[graph[w][i]] = true
				queue.PushQueue(graph[w][i])
			}
		}
	}

}

var found bool

func dfs(graph [8][]int, t int, s int) {
	if t == s {
		return
	}
	//定义visited数组用来保存元素是否访问过
	visited := [8]bool{}

	//定义一个数组用来记录节点从哪里来的
	prev := make([]int, 8)
	for i := 0; i < len(graph); i++ {
		prev[i] = -1
	}

	visited[t] = true
	singDfs(graph, t, s, &visited, prev)
	fmt.Println(prev)
}
func singDfs(graph [8][]int, w int, s int, visited *[8]bool, prev []int) {
	//提行已經找到，防止prev被其他singdfs修改也可以在這裏直接判斷w==s
	if found {
		//fmt.Println(prev)
		return
	}
	visited[w] = true
	for i := 0; i < len(graph[w]) && w != -1; i++ {
		//找到这个定点的一个点
		tmp := graph[w][i]
		//如果这个点没有访问过
		if !visited[tmp] {
			//tmp是从w来的
			prev[tmp] = w
			if tmp == s {
				found = true
				fmt.Println(prev, 114)
				return
			}
			singDfs(graph, tmp, s, visited, prev)
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
