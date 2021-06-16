package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 = [1,2,4], l2 = [1,3,4]
	nodeList1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 4,
	}
	nodeList1.Next = node2
	node2.Next = node3
	nodeList2 := &ListNode{
		Val: 1,
	}
	node5 := &ListNode{
		Val: 3,
	}
	node6 := &ListNode{
		Val: 4,
	}
	nodeList2.Next = node5
	node5.Next = node6
	res := mergeTwoLists(nodeList1, nodeList2)
	show(res)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	prev := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 == nil {
		head.Next = l2
	}
	if l2 == nil {
		head.Next = l1
	}
	return prev.Next
}

func show(head *ListNode) {

	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}

}
