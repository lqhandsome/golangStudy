package main

import "fmt"

func main() {
		first := AddBoy(500)
		showBoy(first)
		playGame(first,20,31)
}

type Boy struct {
	No int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("none")
		return &Boy{}
	}
	for i:= 1 ; i <= num ; i++ {
		boy := &Boy{
			No :i,
		}
		if i == 1 {
			first = boy //头指针
			curBoy = boy
			first.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构成环形链表
		}
	}
	return first
}

func showBoy(first *Boy) {
	//创建一个指针，遍历指针
	curBoy := first
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("小孩编号%v\n",curBoy.No)

		if curBoy.Next == first {
			return
		}
		curBoy = curBoy.Next
	}
}

//留下环形链表按照规定留下一个元素
func playGame(first *Boy,startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("空的链表。没有小孩")
		return
	}
	tail := first
	//让tail指向环形链表的最后一个小孩 删除小孩市使用到
	for  {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	//让first移动到startNo位置
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}
	for {
		for i := 1;i <= countNum - 1;i++  {
			first = first.Next
			tail = tail.Next
		}
		//删除一个小孩

		fmt.Printf("小孩编号为%d 出圈->\n",first.No)
		tail.Next = first.Next
		first  = tail.Next
		//first = first.Next
		//tail.Next = first

		if first.Next == first {
		//if first == tail {
			break
		}
	}
	fmt.Printf("最后出圈的小孩%d",first.No)


}