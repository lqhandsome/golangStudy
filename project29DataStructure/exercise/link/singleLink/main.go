package main

import (
	"errors"
	"fmt"
)

//定义一个节点
type HeroNode struct {
	no       int
	nickName string
	name     string
	next     *HeroNode //表示指向下一个节点
}
func main() {
	head := &HeroNode{}

	name1 := &HeroNode{
		no:       1,
		nickName: "n11",
		name:     "n1",
		next:     nil,
	}
	name2 := &HeroNode{
		no:       2,
		nickName: "n2",
		name:     "n22",
		next:     nil,
	}
	InsertNode(head,name1)
	InsertNode(head,name2)
	ListHeroNode(head)
	DeleteHeroNode(head,1)
	fmt.Println()
	ListHeroNode(head)
}


//插入一个节点
func InsertNode(head *HeroNode,node *HeroNode ) {
	if head.next == nil {
		head.next = node
		return
	}
	tmp := head
	for  {
		tmp = tmp.next
		if tmp.next == nil {
			tmp.next = node
			return
		}
	}
}

//删除一个节点
func DeleteHeroNode(head *HeroNode,no int) (err error) {
	if head.next == nil {
		return  errors.New("链表为空")
	}
	tmp := head.next
	prev := head
	for  {
		if tmp.no == no {
			prev.next = tmp.next
			return nil
		}
		if tmp.next == nil {
			return  errors.New("找不到")
		}
		tmp = tmp.next

	}

}

//显示链表数据
func ListHeroNode(head *HeroNode) {
	if head.next == nil {
		fmt.Println("空链表")
		return
	}
		tmp := head.next
		for {
			fmt.Printf("no=%d,姓名=%v,昵称=%v",tmp.no,tmp.name,tmp.nickName)
			if tmp.next == nil {
				return
			}
			tmp = tmp.next
		}
}


