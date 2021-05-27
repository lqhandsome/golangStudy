package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	graph := [8]*node{}

	node1 := &node{
		val: 1,
	}
	node2 := &node{
		val: 2,
	}
	node3 := &node{
		val: 3,
	}
	node1.next = node2
	node2.next = node3
	//graph[0].next = node1
	graph[0] = &node{
		val: 0,

	}
	graph[0].next = node1
	showList(graph)
}

func showList(graph [8]*node) {
	for i := 0; i < len(graph) && graph[i] != nil; i++ {
		tmp := graph[i].next
		for tmp != nil {
			fmt.Println(tmp.val)
			tmp = tmp.next
		}
	}
}
