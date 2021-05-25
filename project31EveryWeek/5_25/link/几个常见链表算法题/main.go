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
	//ShowLink(head)
	//fmt.Println(detectCycle(head))
	//newHead := removeNthFromEnd(head, 3)
	//newHead := reverseList(head)
	//newHead := reverseListCursion(head)
	ShowLink(head)
	fmt.Println(MidNode(head))
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
		return head
	}
	newHead := reverseListCursion(head.next)
	head.next.next = head
	head.next = nil
	return newHead
}

//3.环的检测查找一个链表中环的位置

func detectCycle(head *List) *List {
	if head == nil {
		return nil
	}
	slow, fast := head.next, head.next.next
	for fast.next.next != nil && fast.next != nil && fast != slow {
		slow = slow.next
		fast = fast.next.next
	}
	p := head
	if fast == slow {
		for p != slow {
			slow = slow.next
			p = p.next
		}
		return p
	} else {
		return nil
	}

}

//4.删除链表的倒数第n个节点
func removeNthFromEnd(head *List, n int) *List {
	if head == nil || head.next == nil {
		return nil
	}
	slow, fast := head, head
	for n != 0 {
		fast = fast.next
		n--
	}
	//fast走到了最后代表要删除的是头结点
	if fast == nil {
		return head.next
	}
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	fmt.Println(slow)
	fmt.Println(fast)
	slow.next = slow.next.next
	return head

}

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
func MidNode(head *List) *List {
	if head == nil || head.next == nil {
		return head
	}
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
