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
	two.next = three
	one.next = two
	head.next = one
	ShowLink(head)
	//newHead := reverseList(head)
	newHead := reverseListCursion(head)
	ShowLink(newHead)
}

//1.反转链表
func reverseList(head *List) (newHead *List) {
	if head.next == nil {
		return
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
func reverseListCursion(head *List) *List {
	if head == nil || head.next == nil {
		return  head
	}
	newHead := reverseListCursion(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

//3.环的检测

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
