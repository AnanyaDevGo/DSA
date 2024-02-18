package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isValidBST(root *TreeNode) bool {
// 	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
// }

func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return isValidBSTHelper(node.Left, min, node.Val) && isValidBSTHelper(node.Right, node.Val, max)
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	isValid := isValidBSTHelper(root, math.MinInt, math.MaxInt)

	if isValid {
		fmt.Println("The binary tree is a valid binary search tree.")
	} else {
		fmt.Println("The binary tree is not a valid binary search tree.")
	}
}
