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
	//bfs(graph,0,6)
	dfs(graph, 0, 6)

}

func bfs(graph [8][]int, s int, t int) {
	if s == t {
		fmt.Println(1111)
		return
	}
	var prev [len(graph)]int
	for i := 0; i < 8; i++ {
		prev[i] = -1
	}

	//保存一个节点是否访问过,访问过为-1
	visited := [len(graph)]bool{}
	//用来保存该节点已经被访问,但是指向的节点还没被访问的节点
	queue := &Queue{
		arr: [8]int{-1, -1, -1, -1, -1, -1, -1},
	}
	queue.PushQueue(s)
	visited[s] = true
	//判断队列中有值
	for queue.front != queue.rear {
		//从队列中取一个顶点,顶尖不用判断是否等于终点,因为等于就不会压入了
		w, _ := queue.PopQueue()
		//遍历这个定点的所有边,如果这个定点的边不符合查找的元素切没有被访问过,就放入队列
		for i := 0; w != -1 && i < len(graph[w]); i++ {
			//一个边
			tmp := graph[w][i]
			if !visited[tmp] {
				//记录tmp是从哪里来的
				prev[tmp] = w
				if tmp == t {
					fmt.Println(prev)
					fmt.Println(visited)
					return
				} else {
					visited[tmp] = true
					//将这个边压入队列
					queue.PushQueue(tmp)
				}

			}
		}
	}
}

var found bool

func dfs(graph [8][]int, s int, t int) {

	var visited = [8]bool{}
	prev := make([]int,8)
	found = false
	if s == t {
		return
	}
	for i := 0; i < 8; i++ {
		prev[i] = -1
	}

	//保存一个节点是否访问过,访问过为-1
	recurDFS(graph, s, t, &visited, prev)
	fmt.Println(prev)
}
func recurDFS(graph [8][]int, w int, t int, visited *[8]bool, prev []int) {
	visited[w]  = true
	if w == t {
		return
	}
	for i := 0; i < len(graph[w]); i++ {
		tmp := graph[w][i]
		if !visited[tmp] {
			prev[tmp] = w
			recurDFS(graph, tmp, t, visited, prev)
		}
	}
}

func showList(graph [8][]int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Print(graph[i][j], " ")
		}
		fmt.Println()
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

func (queue *Queue) ShowQueue() {
	fmt.Println("打印队列")
	if queue.front == queue.rear {
		fmt.Println("队列为空")
		return
	}
	for i := queue.front; i < queue.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, queue.arr[i])
	}
	fmt.Println()
}
