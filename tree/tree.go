package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {
	var root TreeNode
	root = TreeNode{1, nil, nil}
	root.left = &TreeNode{2, nil, nil}
	root.right = &TreeNode{}
	root.right.left = new(TreeNode)
	root.left.right = creatTreeNode(2)
	root.right.left.setValue(4)

	root.traverse()
}

func (node TreeNode) print() {
	fmt.Println(node.value)
}

func (node *TreeNode) setValue(value int) *TreeNode {
	node.value = value
	return node
}

//工厂函数
func creatTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}
//中式遍历
func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
