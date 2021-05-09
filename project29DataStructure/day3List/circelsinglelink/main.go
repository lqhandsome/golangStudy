package main

import "fmt"

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
	//cat2 := &CatNode{
	//	no:   2,
	//	name: "Tom2",
	//}
	//cat3 := &CatNode{
	//	no:   3,
	//	name: "Tom3",
	//}
	InsertCatNode(head,cat1)
	//InsertCatNode(head,cat2)
	//InsertCatNode(head,cat3)
	ListCircleLink(head)
	//head = DeleteCircleLink(head,1)
	head.del(1)
	ListCircleLink(head)

}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		//头猫
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		return
	}

	tmp := head
	for  {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}

	//加入到链表中
	tmp.next = newCatNode
	newCatNode.next = head
}

//遍历环形链表
func ListCircleLink(head *CatNode){
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的链表")
		return
	}

	for {
		fmt.Println("小猫猫",temp,"->")
		if temp.next == head {
			return
		}
		temp = temp.next
	}
}
func (head *CatNode) del(no int) {
	if head.next == nil {
		fmt.Println("链表为空，无法删除")
		return
	}

	if head.next == head  {
		if head.no == no {
			fmt.Println(111)
			head.next = nil
			head.no = 0
			head.name = ""

			return
		}
		return
	}
}
//删除一个
func DeleteCircleLink(head *CatNode,no int) *CatNode {
	if head.next == nil {
		fmt.Println("链表为空，无法删除")
		return head
	}

	if head.next == head  {
		if head.no == no {
			fmt.Println(111)
			//head = &CatNode{}
			head.next = nil
			head.no = 0
			head.name = ""
			return head
		}
		return head
	}

	//tmp := head.next
	//helper := head

	//for {
	//
	//	tmp = tmp.next
	//	helper = helper.next
	//	if tmp.no == no {
	//
	//	}
	//}
	return head
}