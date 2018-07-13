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
	root.left.right = creatTreeNode(5)
	root.right.left.setValue(4)

	root.traverse()

}

//工厂函数
func creatTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func (node *TreeNode) setValue(value int) *TreeNode {
	node.value = value
	return node
}

func (node *TreeNode) traversefunc(f func(*TreeNode))  {

	if node == nil {
		return
	}
	node.left.traversefunc(f)
	f(node)
	node.right.traversefunc(f)
}

func (node *TreeNode) traverse(){
	node.traversefunc(func(n *TreeNode) {
		fmt.Println(n.value)
	})
}