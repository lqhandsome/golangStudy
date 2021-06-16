package main

import (
	"fmt"
)

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
	//four := &List{
	//	5,
	//	nil,
	//}
	//three.next = four
	two.next = three
	one.next = two
	head.next = one
	//four.next = two
	 //p := detectCycle(head)
	 //fmt.Println(detectCycleMap(head))
	//ShowLink(head)
	//newHead := reverseList(head)
	//ShowLink(newHead)
	//newHead2 := reverseListCursion(head)
	//ShowLink(newHead2)
	fmt.Println(MidNode(head))

}

//1.反转链表
func reverseList(head *List) (newHead *List) {
	if head == nil || head.next == nil {
		return head
	}
	var prev  = &List{}
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
		return head
	}
	prev := reverseListCursion(head.next)
	head.next.next = head
	head.next = nil
	return prev
}

//3.环的检测查找一个链表中环的位置
func detectCycle(head *List) *List {
	if head == nil || head.next == nil {
		return nil
	}
	slow , fast := head.next,head.next.next
	for slow != nil {
		if slow == fast {
			fmt.Println(slow.val)
			p := head
			for p != slow {
				p = p.next
				slow = slow.next
			}
			return p
		}
		slow = slow.next
		fast = fast.next.next
	}
	return nil
}

func detectCycleMap(head *List) *List {
	if head == nil || head.next == nil {
		return nil
	}
	seen := map[*List]int{}
	tmp := head
	for tmp != nil {
		_,ok := seen[tmp]
		if ok {
			return tmp
		} else {
			seen[tmp] = 1
		}
		tmp = tmp.next
	}
	return nil
}

//4.删除链表的倒数第n个节点
func removeNthFromEnd(head *List, n int) *List {
	return head
}

//显示链表所有数据
func ShowLink(head *List) {
	tmp := head
	for tmp != nil {
		fmt.Print(tmp.val,"->")
		tmp = tmp.next
	}
	fmt.Println("nil")
}

//求链表的中间节点
func MidNode(head *List) *List {
	if head == nil || head.next == nil {
		return head
	}
	slow,fast := head,head
	for fast.next.next != nil{
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
