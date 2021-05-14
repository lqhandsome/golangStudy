package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int    //最大存放个数
	Top    int    //表示栈顶
	arr    [5]int //存放数据
}

//入栈
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("栈满")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈为空")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Println(this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (err error, val int) {
	if this.Top == -1 {
		return errors.New("栈为空"), -1
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

//判断一个字符是不是运算符
func IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

//运算
func Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num1 * num2
	case 43:
		res = num1 + num2
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符有问题")
		return -1
	}
	return res
}

//运算符优先级 （0）+ - （1）* / （2） （）
func Priority(oper int) int {
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 0
	}
	fmt.Println("ERR")
	return -1

}
func main() {
	//数字栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	//符合栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "8+2*6/3-2*2*2+5+1"
	// index = 0
	index := 0
	var num1, num2, oper, result int
	for index < len(exp) {
		tmp := int(exp[index]) //是一个字符串
		//fmt.Println(cn)
		if IsOper(tmp) {
			//说明是符号

			//符号栈为空
			if operStack.Top == -1 {
				operStack.Push(tmp)
			} else {
				if Priority(operStack.arr[operStack.Top]) >= Priority(tmp) {
					_, num1 = numStack.Pop()
					_, num2 = numStack.Pop()
					_, oper = operStack.Pop()
					result = Cal(num1, num2, oper)
					numStack.Push(result)
					operStack.Push(tmp)
				} else {
					operStack.Push(tmp)
				}
			}
		} else {
			//是数字
			val,_ := strconv.ParseInt(exp[index:index+1],10,64)
			//fmt.Printf("%T%v\n",val,val)
			numStack.Push(int(val))
		}
		index++
		if index == len(exp) {
			break
		}
	}
	numStack.List()
	operStack.List()
	fmt.Println("-------------")
	for operStack.Top != -1  {
		_, num1 = numStack.Pop()
		_, num2 = numStack.Pop()
		_, oper = operStack.Pop()
		fmt.Println(num1,num2,oper)
		result = Cal(num1, num2, oper)
		numStack.Push(result)
	}
	_, res := numStack.Pop()
	fmt.Println(exp,"=",res)
}
