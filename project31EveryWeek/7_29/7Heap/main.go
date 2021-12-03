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

}

//往堆里插入一个元素从下往上堆化
func (this *Heap) insert(val int) error {

}

//删除堆顶元素
func (this *Heap) deleteHeap() error {

	return nil
}
