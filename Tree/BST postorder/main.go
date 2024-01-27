package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func BuildBSTFromPostorder(postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Value: postorder[len(postorder)-1]}
	i := 0

	for i < len(postorder)-1 && postorder[i] < root.Value {
		i++
	}

	root.Left = BuildBSTFromPostorder(postorder[:i])
	root.Right = BuildBSTFromPostorder(postorder[i : len(postorder)-1])

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
	postorder := []int{1, 7, 5, 12, 10, 8}


	root := BuildBSTFromPostorder(postorder)


	fmt.Println("In-order traversal of the constructed BST:")
	InorderTraversal(root)
	fmt.Println()
}
