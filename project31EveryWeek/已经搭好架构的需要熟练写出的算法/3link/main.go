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

}

//显示链表所有数据
func ShowLink(head *Link) {


}

//删除一个节点
func Delete(head *Link, no int) (err error) {

}
