package main

import "fmt"

type Link struct {
	val  int
	next *Link
}

func main() {
	head := &Link{}
	node1 := &Link{
		1,
		nil,
	}
	node2 := &Link{
		2,
		nil,
	}
	node3 := &Link{
		3,
		nil,
	}
	head.insert(node1)
	head.insert(node2)
	head.insert(node3)
	head.show()
	head.delete(1)
	head.show()
}

func (this *Link) insert(node *Link) {
	if this.val == 0 {
		this.val = node.val
		return
	}
	tmp := this
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
}
func (this *Link) delete(val int) {

	tmp,prev := this.next,this
	//如果要删除的是头节点
	if prev.val == val {
		if prev.next != nil {
			prev.val = prev.next.val
			prev.next = prev.next.next
		} else {
			prev.val = 0
			prev.next = nil
		}

		return
	}
	for tmp != nil {
		if tmp.val == val {
			prev.next = tmp.next
			return
		}
		tmp,prev = tmp.next,prev.next
	}
}

func (this *Link) show() {
	tmp := this
	for tmp != nil && tmp.val != 0{
		fmt.Print(tmp.val,"->")
		tmp = tmp.next
	}
	fmt.Println()
}
