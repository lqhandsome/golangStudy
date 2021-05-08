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
	prev     *HeroNode //表示指向上一个节点
}

func main() {
	//创建一个头节点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "武松",
		nickName: "沙和尚",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "鲁智深",
		nickName: "沙和尚",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickName: "沙和尚",
	}
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero4)
	//DeleteHeroNode(head,3)
	//fmt.Println(head.next.next)
	ListHeroNode(head)
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	tmp := head
	for {
		if tmp.next == nil {
			tmp.next = newHeroNode
			newHeroNode.prev = tmp
			return
		}
		tmp = tmp.next
	}
}

/**
顺序添加节点 按照no的编号从小到大插入
*/
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode)  {
	tmp := head
	for {
		if tmp.next == nil {
			newHeroNode.next = tmp.next
			tmp.next = newHeroNode
			return
		} else if tmp.next.no > newHeroNode.no {
			newHeroNode.next = tmp.next
			tmp.next = newHeroNode
			return
		} else if tmp.next.no == newHeroNode.no {
			fmt.Println("存在相同的no",newHeroNode.no)
			return
		}
		tmp = tmp.next
	}
}

func ListHeroNode(head *HeroNode) {
	tmp := head
	for {
		if tmp.next == nil {
			return
		}
		fmt.Println(tmp)
		tmp = tmp.next
	}
}

func DeleteHeroNode(head *HeroNode,no int) (err error) {
	tmp := head
	for  {
		if tmp.next.no == no {
			tmp.next = tmp.next.next
			return
		}

		if tmp.next == nil {
			return errors.New("没有找到对应的元素")
		}
		tmp = tmp.next

	}
	return errors.New("没有找到对应的元素")
}
