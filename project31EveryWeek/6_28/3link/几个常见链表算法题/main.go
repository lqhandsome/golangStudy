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
	three.next = two
	two.next = three
	one.next = two
	head.next = one
	//ShowLink(head)
	//newHead := reverseList(head)
	//newHead := removeNthFromEnd(head,3)
	//newHead := MidNode(head)
	//ShowLink(newHead)
	//fmt.Println(detectCycleMap(head))
	fmt.Println(detectCycle(head))

}

//1.反转链表
func reverseList(head *List) (newHead *List) {
	if head == nil{
		 return nil
	}
	prev := new(List)
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
	prev := reverseList(head.next)
	head.next.next = head
	head.next = nil
	return prev
}

//3.环的检测查找一个链表中环的位置

func detectCycle(head *List) *List {
	if head == nil {
		return nil
	}
	slow,fast := head.next,head.next.next
	for fast != nil && fast.next != nil {
		//快慢指针相遇了，表示有环
		if slow == fast {
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
	m := make(map[*List]bool,10)
	tmp := head
	for tmp != nil {
		if _,ok := m[tmp];ok {
			return tmp
		}
		m[tmp] = true
		tmp = tmp.next
	}
	return nil
}

//4.删除链表的倒数第n个节点
func removeNthFromEnd(head *List, n int) *List {
	//定义快慢指针
	slow,fast := head,head
	//先让快指针走n步
	for n > 0 {
		fast = fast.next
		n--
	}
	//删除的是头节点
	if fast == nil {
		return head.next
	}
	for fast != nil && fast.next != nil{
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	return head
}

//显示链表所有数据
func ShowLink(head *List) {
	tmp := head
	for tmp != nil {
		fmt.Print(tmp.val,"->")
		tmp = tmp.next
	}
	fmt.Println()
}

//求链表的中间节点
func MidNode(head *List) *List {
	//定义快慢指针
	slow,fast := head,head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
