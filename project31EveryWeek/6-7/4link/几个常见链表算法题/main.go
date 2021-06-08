package main

import "fmt"

type List struct {
	val  int
	next *List
}

func main() {
	head := &List{
		1,
		nil,
	}
	one := &List{
		2,
		nil,
	}

	two := &List{
		3,
		nil,
	}
	three := &List{
		4,
		nil,
	}
	//three.next = two
	two.next = three
	one.next = two
	head.next = one
	ShowLink(head)
	//fmt.Println(detectCycle(head))
	//newHead := removeNthFromEnd(head, 3)
	//newHead := reverseList(head)
	newHead := reverseListCursion(head)
	ShowLink(newHead)
	//fmt.Println(MidNode(head))
}

//1.反转链表
func reverseList(head *List) (newHead *List) {
	//只有一个节点
	if head.next == nil {
		return head
	}
	prev := &List{}
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

//2.反转链表递归写法
func reverseListCursion(head *List) (newHead *List) {
	if head == nil || head.next == nil {
		return head
	}
	//用来接收已经翻转好的数据
	prev := reverseListCursion(head.next)
	head.next.next = head
	head.next = nil
	return prev
}
//3.环的检测查找一个链表中环的位置

//4.删除链表的倒数第n个节点

//显示链表所有数据
func ShowLink(head *List) {
	if head == nil || head.next == nil {
		return
	}
	tmp := head
	for {
		fmt.Print(tmp.val, "->")
		tmp = tmp.next
		if tmp == nil {
			fmt.Println("完成")
			return
		}
	}

}

//求链表的中间节点
