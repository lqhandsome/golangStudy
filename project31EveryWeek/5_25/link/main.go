package main

import (
	"errors"
	"fmt"
)

type Link struct {
	val  int
	next *Link
}

func main() {
	head := &Link{}

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

//插入一个节点
func Insert(head *Link, insertVal *Link) (err error) {
	if head.next == nil {
		head.next = insertVal
		return err
	}
	tmp := head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = insertVal
	return err
}

//显示链表所有数据
func ShowLink(head *Link) {
	if head.next == nil {
		return
	}
	tmp := head
	for  {
		fmt.Print(tmp.val, "->")
		tmp = tmp.next
		if tmp == nil {
			fmt.Println("完成")
			return
		}
	}

}

//删除一个节点
func Delete(head *Link, no int) (err error) {
	if head.next == nil {
		return errors.New("空队列，无法删除")
	}
	prev := head
	tmp := head.next
	for tmp != nil {
		if tmp.val == no {
			prev.next = tmp.next
			return errors.New("删除成功")
		}
		prev = prev.next
		tmp = tmp.next
	}
	return errors.New("找不到")
}
