package main

import (
	"errors"
	"fmt"
)

type Link struct {
	val  int
	next *Link
}

//普通链表
func main() {
	head := &Link{val: -1}

	Insert(head, &Link{
		1,
		nil,
	})
	Insert(head, &Link{
		2,
		nil,
	})
	Insert(head, &Link{
		3,
		nil,
	})
	ShowLink(head)
	fmt.Println(Delete(head,2))
	ShowLink(head)
}

// Insert 插入一个节点
func Insert(head *Link, insertVal *Link) (err error) {
	//如果头节点为空
	if head == nil {
		fmt.Println("empty head")
		return errors.New("empty head")
	}
	tmp := head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = insertVal
	return nil
}

// ShowLink 显示链表所有数据
func ShowLink(head *Link) {
	tmp := head
	for tmp != nil {
		fmt.Printf("node=%v\n",tmp.val)
		tmp = tmp.next
	}
}

//删除一个节点
func Delete(head *Link, no int) (err error) {
	if head == nil{
		fmt.Println("链表为空！")
		return nil
	}
	tmp := head
	for tmp != nil {
		if tmp.next != nil && tmp.next.val == no {
			tmp.next = tmp.next.next
		}
		tmp = tmp.next
	}
	return nil
}
