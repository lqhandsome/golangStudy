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
}

//插入一个节点
func (head *Link)Insert( insertVal *Link) (err error) {

}

//显示链表所有数据
func ShowLink(head *Link) {

}

//删除一个节点
func (head *Link)Delete( no int) (newHead *Link,err error) {

}

func testDefer(t int) int {
	defer  func() {
		//x += 100
		t += 100
	}()
	return  t
}