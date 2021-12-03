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

}

//1.反转链表
func reverseList(head *List) (newHead *List) {

}

//2.反转链表递归写法
func reverseListCursion(head *List) *List {

}

//3.环的检测查找一个链表中环的位置

func detectCycle(head *List) *List {


}

//4.删除链表的倒数第n个节点
func removeNthFromEnd(head *List, n int) *List {

}

//显示链表所有数据
func ShowLink(head *List) {

}

//求链表的中间节点
func MidNode(head *List) *List {
}
