package main

import (
	"errors"
	"fmt"
)

//定义猫的结构体节点
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func main() {

	//初始化一下环形头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "Tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "Tom3",
	}
	InsertCatNode(head, cat1)
	//fmt.Println(head)
	//return
	InsertCatNode(head,cat2)
	InsertCatNode(head,cat3)
	head, err := DeleteCircleLink(head,3)

	if err != nil {
		fmt.Println("删除失败" + err.Error())
	}
	ListCircleLink(head)

}

//插入一个node
func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//头节点为空
	if head.next == nil {
		head.name = newCatNode.name
		head.no = newCatNode.no
		head.next = head
		return
	}
	tmp := head.next
	for {
		if tmp.next == head {
			tmp.next = newCatNode
			newCatNode.next = head
			return
		}
		tmp = tmp.next
	}
}

//遍历环形链表
func ListCircleLink(head *CatNode) {
	tmp := head
	for {

		fmt.Println(tmp)
		tmp = tmp.next
		if tmp == head {
			return
		}

	}
}

//删除一个
func DeleteCircleLink(head *CatNode, no int) (first *CatNode,err error) {
	//空数据
	if head.next == nil {
		return head,errors.New("空链表无法删除")
	}



	tmp := head
	tail := head

	//将指针指向尾部
	for  {
		if tail.next == head {
			break
		}
		tail = tail.next
	}
	//删除的是头节点的数据
	if head.no == no {
		tail.next = head.next
		return tail.next,errors.New("删除成功")
	}
	for  {
		if tmp.no == no {
			tail.next = tmp.next
			return head, errors.New("删除成功")
		}
		tmp = tmp.next
		tail = tail.next
	}
	//删除节点
}
