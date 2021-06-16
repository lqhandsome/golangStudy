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
	ShowLink(head)
	head.Insert(&Link{
		val: 1,
	})
	head.Insert(&Link{
		val: 2,
	})
	head.Insert(&Link{
		val: 3,
	})
	ShowLink(head)
	Head,_ := head.Delete(3)
	ShowLink(Head)
}

//插入一个节点
func (head *Link)Insert( insertVal *Link) (err error) {
	//如果头节点为空
	if head.val == -1 {
		head.val = insertVal.val
		head.next = head
		return nil
	}
	tmp := head
	for tmp.next != head {
		tmp = tmp.next
	}
	tmp.next = insertVal
	tmp.next.next = head
	return err
}

//显示链表所有数据
func ShowLink(head *Link) {
	tmp := head
	for head.next != nil {
		fmt.Print(tmp.val,"->")
		tmp = tmp.next
		if tmp == head {
			break
		}
	}
	fmt.Println()
}

//删除一个节点
func (head *Link)Delete( no int) (newHead *Link,err error) {

	if head.next == nil {
		fmt.Println("队列为空")
		return head,errors.New("队列为空")
	}
	tmp := head.next
	//定义一个节点指向环形队列的尾部,用来删除头节点
	tail := head
	for tail.next != head {
		tail = tail.next
	}
	//删除的是头节点
	if head.val == no {
		//头节点只有一个值
		if head.next == head {
			return &Link{}, err
		}
		tail.next = head.next
		return tail.next,nil
	}

	for tmp != head {
		if tmp.val == no {
			tail.next.next = tmp.next
		}
		tail =tail.next
		tmp = tmp.next
	}
	return head,err
}

func testDefer(t int) int {
	 defer  func() {
		//x += 100
		t += 100
	}()
	return  t
}