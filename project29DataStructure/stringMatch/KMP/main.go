package main

import (
	"fmt"
)

type Stack struct {
	arr []int
	top int
}

func inset(c chan int) {
	var count int
	for  {
		count++
		//time.Sleep(time.Second)
		if count > 10 {
			close(c)
			return
		}
		c<-1
	}
}

func main() {
	//str := "BABBBABABAAA"
	//findStr := "ABABA"
	//next := []int{0, 1, 1, 2, 3}
	//fmt.Println(KMP(str, findStr, next))

	c := make(chan int,1)
	go inset(c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(c)
	return
	s := Stack{
		make([]int, 5),
		-1,
	}
	s.Push(4)
	s.Push(2)
	s.Push(1)
	s.Push(3)
	fmt.Println(s.sortStack().arr)
}
//将一个堆栈中的数据排列
func (this Stack) sortStack() Stack {
	tmpTop := this.top
	res :=Stack{
		make([]int, 5),
		-1,
	}
	var newStack Stack = Stack{
		make([]int, 5),
		-1,
	}
	for this.top != -1 {
		//用来保存一次最小的数
		var tmp int

		//tmpTop--
		//找到最小的数
		for this.top != -1 {
			if tmp < this.arr[this.top] {
				tmp = this.arr[this.top]
			}
			newStack.Push(this.Pop())
		}
		//重新吧数据放入this
		for i := 0; i <= tmpTop; i++ {
			popVal := newStack.Pop()
			if popVal == tmp {
					continue
			} else {
				this.Push(popVal)
			}

		}
		newStack.Push(tmp)
		res.Push(tmp)
		tmpTop--
		if tmpTop == -1 {
			 return newStack
		}

	}
	return Stack{}
}

func KMP(str string, findStr string, next []int) int {
	//判断在模式串中的位置，当j == len(findStr) 代表找到
	j := 0
	for i := 0; i < len(str); i++ {
		//除第一个发生不匹配的情况
		for j > 0 && str[i] != findStr[j] {
			j = next[j] - 1
		}
		if str[i] == findStr[j] {
			j++
		}
		if j == len(findStr) {
			return i - len(findStr) + 1
		}
	}
	return -1
}

func (this *Stack) Pop() int {
	if this.top == -1 {
		fmt.Println("栈空")
		return -1
	}
	val := this.arr[this.top]
	this.arr[this.top] = 0
	this.top--
	return val
}

func (this *Stack) Push(val int) bool {
	this.top++
	this.arr[this.top] = val
	return true
}
