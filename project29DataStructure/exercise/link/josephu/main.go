package main

import "fmt"

type Boy struct {
	No int
	Next *Boy
}
//新增boy
func AddBoy(num int) *Boy {
	boy := &Boy{}
	head := boy
	for i := 1; i <= num; i++ {
		if i == 1 {
			boy.Next = boy
			boy.No = i
		} else {
			curBoy := &Boy{
				No: i,
				Next: head,
			}
			boy.Next = curBoy
		}
		boy = boy.Next
	}
	return head

}

func showBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	tmp := first
	for  {
		fmt.Printf("小孩编号%d\n",tmp.No)
		tmp = tmp.Next
		if  tmp == first {
			return
		}
	}
}
//留下环形链表按照规定留下一个元素
func playGame(first *Boy,startNo int, countNum int) {
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	tmp := first

	//设置尾结点
	tailNode := first
	for  {
		if tailNode.Next == first {
			break
		}
		tailNode = tailNode.Next
	}
	fmt.Println("尾结点",tailNode.No)
	//确定起点的位置
	for i:= 1; i <= startNo - 1 ; i++{
		tmp = tmp.Next
		tailNode = tailNode.Next
		fmt.Println("起点编号",tmp.No)
	}
	for {

		//遍历删除节点
		for i := 1;i < countNum ;i++ {
			tailNode = tailNode.Next
			tmp = tmp.Next
		}
		fmt.Println("小孩被删除了",tmp.No)
		tailNode.Next = tmp.Next
		tmp = tailNode.Next

		if tmp.Next == tmp {
			break
		}
	}
	fmt.Println("最后一个小孩是",tmp.No)
}
func main() {
	boy := AddBoy(5)
	showBoy(boy)
	playGame(boy,3,3)
}
