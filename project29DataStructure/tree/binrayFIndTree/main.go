package main

import "fmt"

type tree struct {
	val int
	left *tree
	right *tree
}
func main() {
	//tree := resTree()
	//midRangeTree(tree)
	//fmt.Println(findRecursion(tree,6))
	//fmt.Println(find(tree,3))
	t := &tree{
		val: 3,
	}
	insert(t,2)
	insert(t,5)
	midRangeTree(t)
	fmt.Println(t)
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

//二叉查找树的插入操作
/*
如果要插入的数据比节点的数据大，并且节点的右子树为空，就将新数据直接插到右子节点的位置；
如果不为空，就再递归遍历右子树，查找插入位置。
同理，如果要插入的数据比节点数值小，并且节点的左子树为空，就将新数据插入到左子节点的位置；
如果不为空，就再递归遍历左子树，查找插入位置
 */
func insert(treee *tree,n int) {

	if treee.val == 0 {
		treee.val = n
		return
	}
	if n > treee.val {
		if treee.right == nil {
			treee.right = &tree{
				val: n,
			}
			return
		}
		insert(treee.right,n)
	} else {
		if treee.left == nil {
			treee.left = &tree{
				val: n,
			}
			return
		}
		insert(treee.left,n)
	}
}
func delete(treee *tree,n int)  {
	if treee == nil {
		return
	}
	p := treee
	pp := &tree{}
	//找到要删除的节点
	for p != nil && p.val != n {
		//父节点
		pp = p
		if n > p.val {
			p = p.right
		} else {
			p = p.left
		}
	}
	//没有找到
	if p == nil {
		return
	}
	//删除节点是叶子节点或者只有一个节点
	if p.left == nil || p.right == nil {

		//没有子节点
		if p.left == nil && p.right == nil {
			if pp.left == p {
				pp.left = nil
			} else {
				pp.right = nil
			}
		}
		//只有右子节点
		if p.left == nil {
			child := p.right
			if pp.left == p {
				pp.left = child
			} else {
				pp.right = child
			}
		}

		//只有左子节点
		if p.right == nil {
			child := p.left
			if pp.left == p {
				pp.left = child
			} else {
				pp.right = child
			}
		}
	}

	//删除的节点有两个节点


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

