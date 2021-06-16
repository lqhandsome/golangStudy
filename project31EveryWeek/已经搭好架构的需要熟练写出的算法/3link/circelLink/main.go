package main

import (
	"errors"
	"fmt"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

//循环链表
func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		1,
		"小花",
		nil,
	}
	cat2 := &CatNode{
		2,
		"小米",
		nil,
	}
	cat3 := &CatNode{
		3,
		"小玉",
		nil,
	}

	cat4 := &CatNode{
		4,
		"小4",
		nil,
	}
}

//环形链表加入一个节点
func insertNode(head *CatNode, node *CatNode) error {

}

//删除一个节点
func delete(head *CatNode, no int) (newHeadNode *CatNode, err error) {


}

//遍历链表
func ShowCircelLink(head *CatNode) {

}
