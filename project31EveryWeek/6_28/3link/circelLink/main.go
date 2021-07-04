package main

import (
	"fmt"
)

type Link struct {
	val  int
	next *Link
}

//环形链表
func main() {

	head := &Link{
		val: -1,
	}
	node1 := &Link{
		val: 1,
	}
	node2 := &Link{
		val: 2,
	}
	node3 := &Link{
		val: 3,
	}
	head.Insert(node1)
	head.Insert(node2)
	head.Insert(node3)
	ShowLink(head)
	head,_ = head.Delete(2)
	ShowLink(head)
}

//插入一个节点
func (head *Link)Insert( insertVal *Link) (err error) {
	//头节点为空
	if head.val == -1 {
		head.val = insertVal.val
		head.next = head
		return err
	}
	tmp := head
	for tmp.next != head {
		tmp = tmp.next
	}
	tmp.next = insertVal
	insertVal.next = head
	return err
}

//显示链表所有数据
func ShowLink(head *Link) {
	if head.val == -1 {
		fmt.Println("空链表")
		return
	}
	tmp := head
	for true {
		fmt.Print(tmp.val,"->")
		tmp = tmp.next
		if tmp == head{
			 break
		}
	}
}

//删除一个节点
func (head *Link)Delete( no int) (newHead *Link,err error) {
	if head.val == -1 {
		fmt.Println("无法删除空节点")
		return head, err
	}

	//设置一个tail指向头节点，用来删除头节点使用
	tail := head
	for tail.next != head {
		tail = tail.next
	}
	fmt.Println("tail",tail)
	//如果删除的是头节点
	if head.val == no {
		//只有一个头节点
		if head.next == head {
			head.val = -1
			head.next = nil
			return head, err
		}
		tail.next = head.next
		fmt.Println(87,tail.val,tail.next.val)
		return tail.next, err
	}
	tmp := head
	for tmp.val != no {
		tail = tail.next
		tmp = tmp.next
	}
	tail.next = tmp.next
	return head, err
}
