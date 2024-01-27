package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBSTFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: preorder[0]}
	i := 1

	for i < len(preorder) && preorder[i] < root.Value {
		i++
	}

	root.Left = BuildBSTFromPreorder(preorder[1:i])
	root.Right = BuildBSTFromPreorder(preorder[i:])

	return root
}

func InorderTraversal(root *TreeNode) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Printf("%d ", root.Value)
		InorderTraversal(root.Right)
	}
}

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}

	root := BuildBSTFromPreorder(preorder)

	fmt.Println("In-order traversal of the constructed BST:")
	InorderTraversal(root)
	fmt.Println()
}
