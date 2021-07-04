package main

import (
	"errors"
	"fmt"
)

type Heap struct {
	arr   [10]int //储存元素
	n     int     //总共可以储存的元素个数
	count int     //已经存放的输了
}
//对于完全二叉树来说，下标从 2n​+1 到 n 的节点都是叶子节点。
/*
查找前 K 大数据呢？
比小顶堆最小的都小.肯定不符合topK
我们可以维护一个大小为 K 的小顶堆，顺序遍历数组，从数组中取出数据与堆顶元素比较。如果比堆顶元素大，我们就把堆顶元素删除，并且将这个元素插入到堆中；
如果比堆顶元素小，则不做处理，继续遍历数组。这样等数组中的数据都遍历完之后，堆中的数据就是前 K 大数据了。
 */
func main() {
	heap := &Heap{
		n:     9,
		count: 0,
	}
	heap.insert(3)
	heap.insert(1)
	heap.insert(5)
	heap.insert(7)
	fmt.Println(heap.arr)
	heap.deleteHeap()
	fmt.Println(heap.arr)
}

//往堆里插入一个元素从下往上堆化
func (this *Heap) insert(val int) error {
	if this.count >= this.n {
		return errors.New("堆已满")
	}
	this.count++
	//将元素放入数组
	this.arr[this.count] = val
	num := this.count
	for num > 0 {
		if this.arr[num / 2] > this.arr[num]{
			this.arr[num / 2 ] ,this.arr[num] = this.arr[num], this.arr[num / 2]
		}
		num = num / 2
	}
	return nil
}

//删除堆顶元素
func (this *Heap) deleteHeap() error {
	if this.count < 1 {
		fmt.Println("堆为空")
		return errors.New("堆为空")
	}
	this.arr[1] = this.arr[this.count]
	this.count--
	num := 1
	for num * 2 <= this.count {
		if this.arr[num] > this.arr[num * 2] {
			this.arr[num],this.arr[num * 2] = this.arr[num * 2],this.arr[num]
			num = num * 2
		}
		if this.arr[num / 2] > this.arr[num + 1] {
			this.arr[num / 2],this.arr[num  + 1] = this.arr[num  + 1] , this.arr[num / 2 ]
			num = num  + 1
		}
		fmt.Println(1)
	}
	return nil
}
