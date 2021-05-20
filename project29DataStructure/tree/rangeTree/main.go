package main

import "fmt"

type tree struct {
	val int
	left *tree
	right *tree
}

func main() {
	//普通二叉树的遍历（链表）
	tree := resTree()
	fmt.Println("前序遍历")
	preRangeTree(tree)
	fmt.Println("中序")
	midRangeTree(tree)
	fmt.Println("后序遍历")
	postRangeTree(tree)
}

func preRangeTree(tree *tree) {
	if tree == nil {
		return
	}
	fmt.Println(tree.val)
	preRangeTree(tree.left)
	preRangeTree(tree.right)
}

func midRangeTree(tree *tree) {
	if tree == nil {
		return
	}
	midRangeTree(tree.left)
	fmt.Println(tree.val)
	midRangeTree(tree.right)
}

func postRangeTree(tree *tree) {
	if tree == nil {
		return
	}
	postRangeTree(tree.left)
	postRangeTree(tree.right)
	fmt.Println(tree.val)
}

func resTree() *tree {
	t1 := &tree{
		val: 1,
	}
	t2 := &tree{
		val: 2,
	}
	t3 := &tree{
		val: 3,
	}
	t4 := &tree{
		val: 4,
	}
	t5 := &tree{
		val: 5,
	}
	t1.left = t2
	t1.right = t3
	t2.left = t4
	t2.right = t5
	return  t1
}
