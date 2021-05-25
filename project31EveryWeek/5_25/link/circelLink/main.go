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
	insertNode(head, cat1)
	insertNode(head, cat2)
	insertNode(head, cat3)
	insertNode(head, cat4)
	ShowCircelLink(head)
	head, _ = delete(head, 3)
	head, _ = delete(head, 4)
	head, _ = delete(head, 1)
	head, err := delete(head, 2)
	if err != nil {
		fmt.Println(err)
	}
	ShowCircelLink(head)
}

//环形链表加入一个节点
func insertNode(head *CatNode, node *CatNode) error {
	//头节点下没有数据
	if head.next == nil {
		head.next = head
		head.no = node.no
		head.name = node.name
		return nil
	}
	tmp := head
	for tmp.next != head {
		tmp = tmp.next
	}
	tmp.next = node
	node.next = head
	return nil

}

//删除一个节点
func delete(head *CatNode, no int) (newHeadNode *CatNode, err error) {
	if head.next == nil {
		return head, errors.New("空节点")
	}
	tail := head
	//将尾部指针指向链表头节点
	for tail.next != head {
		tail = tail.next
	}
	//如果删除的头节点
	if head.no == no {
		if head.next == head {
			return &CatNode{}, err
		}
		tail.next = head.next
		return tail.next, nil
	}
	tmp := head.next
	prev := head
	for tmp != head {
		if tmp.no == no {
			prev.next = tmp.next
			return head, err
		}
		tmp = tmp.next
		prev = prev.next
	}
	return nil, err

}

//遍历链表
func ShowCircelLink(head *CatNode) {
	if head.next == nil || head == nil{
		fmt.Println("错误，链表为空")
		return
	}
	tmp := head
	for {
		fmt.Printf("小猫编号%v名字%s->", tmp.no, tmp.name)
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
	fmt.Println("结束")
	return
}
