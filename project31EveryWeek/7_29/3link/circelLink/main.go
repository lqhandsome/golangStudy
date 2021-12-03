package main

import (
	"errors"
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
	fmt.Println(head)
}

//插入一个节点
func (head *Link) Insert(insertVal *Link) (err error) {
	//如果头节点为空
	if head.val == -1 {
		head.val = insertVal.val
		return err
	}
	tmp := head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = insertVal
	insertVal.next = tmp
	return err
}

//显示链表所有数据
func ShowLink(head *Link) {
	if head.val == -1 {
		fmt.Println("空链表")
		return
	}
	tmp := head
	for tmp != nil {
		fmt.Printf("%v->", tmp.val)
		tmp = tmp.next
		if tmp == head {
			break
		}
	}
}

//删除一个节点
func (head *Link) Delete(no int) (newHead *Link, err error) {
	//如果节点为空
	if head.val == -1 {
		fmt.Println("节点为空无法删除")
		return nil, errors.New("节点为空无法删除")
	}

	//如果删除的的是头节点
	return
}
