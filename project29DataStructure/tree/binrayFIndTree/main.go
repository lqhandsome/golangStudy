package main

import "fmt"

type tree struct {
	val int
	left *tree
	right *tree
}
func main() {
	tree := resTree()
	midRangeTree(tree)
	fmt.Println(findRecursion(tree,6))
	fmt.Println(find(tree,3))
}

func resTree() *tree {
	t1 := &tree{
		val: 4,
	}
	t2 := &tree{
		val: 2,
	}
	t3 := &tree{
		val: 5,
	}
	t4 := &tree{
		val: 1,
	}
	t5 := &tree{
		val: 3,
	}
	t1.left = t2
	t1.right = t3
	t2.left = t4
	t2.right = t5
	return  t1
}

func insert(tree *tree) {

}

//二叉查找树的递归查找
func findRecursion(tree *tree,n int) (val int) {
	if tree == nil {
		 return  -1
	}
	if tree.val == n {
		return n
	} else if tree.val < n {
		val = findRecursion(tree.right,n)
	} else {
		val = findRecursion(tree.left,n)
	}
	return val
}

//二叉查找树的迭代查找
func find(tree *tree,n int) (val int) {
	for tree != nil {
		if tree.val == n {
			fmt.Println("找到了")
			return n
		} else if tree.val < n {
			tree = tree.right
		} else {
			tree = tree.left
		}
	}
	fmt.Println("不好意思没找到")
	return  -1
}
//二叉查找树的中序遍历是排序
func midRangeTree(tree *tree) {
	if tree == nil {
		return
	}
	midRangeTree(tree.left)
	fmt.Println(tree.val)
	midRangeTree(tree.right)
}

