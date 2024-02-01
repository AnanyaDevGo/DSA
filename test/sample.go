package main

import (
	"fmt"
	"math"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key}
	}
	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else {
		root.Right = insert(root.Right, key)
	}
	return root
}
func search(root *Node, key int) *Node {
	if root == nil || root.Key == key {
		return root
	}
	if key < root.Key {
		return search(root.Left, key)
	}
	return search(root.Right, key)
}

func preorderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d", root.Key)
		preorderTraversal(root.Left)
		preorderTraversal(root.Right)
	}
}

func findMin(root *Node) *Node {
	current := root
	if current.Left != nil {
		current = current.Left
	}
	return current
}

func findMax(root *Node) *Node {
	current := root
	if current.Right != nil {
		current = current.Right
	}
	return current
}

func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d", root.Key)
		inorderTraversal(root.Right)
	}
}

func postorderTraversal(root *Node) {
	if root != nil {
		postorderTraversal(root.Left)
		postorderTraversal(root.Right)
		fmt.Printf("%d", root.Key)
	}
}

func closest(root *Node, target int) int {
	closestVal := math.Inf(1)
	for root != nil {
		if diff := float64(root.Key - target); math.Abs(diff) < math.Abs(closestVal-float64(target)) {
			closestVal = float64(root.Key)
		}
		if target < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return int(closestVal)
}

func isValidBinaryTree(root *Node, minVal, maxVal int) bool {
	if root == nil {
		return true
	}
	if !(minVal < root.Key && root.Key < maxVal) {
		return false
	}
	return isValidBinaryTree(root.Left, minVal, root.Key) && isValidBinaryTree(root.Right, root.Key, maxVal)
}

func main() {
	var root *Node
	keys := []int{8, 10, 9, 11, 7, 6}

	for _, key := range keys {
		root = insert(root, key)
	}

	fmt.Println("In-order traversal:")
	inorderTraversal(root)
	fmt.Println("\nSearch for 9 :")
	result := search(root, 9)
	if result != nil {
		fmt.Println(result.Key)
	} else {
		fmt.Println("Not found")
	}
	fmt.Println("\nMinimum Number :")
	minNode := findMin(root)
	if minNode != nil {
		fmt.Println(minNode.Key)
	} else {
		fmt.Println("Tree is empty")
	}
	fmt.Println("\nMaximum Number :")
	maxNode := findMax(root)
	if minNode != nil {
		fmt.Println(maxNode.Key)
	} else {
		fmt.Println("Tree is empty")
	}
	fmt.Println("pre-order traversal:")
	preorderTraversal(root)
	fmt.Println("\nPost-order traversal:")
	postorderTraversal(root)

	fmt.Println("\nClosest value to 9:")
	fmt.Println(closest(root, 9))
	fmt.Println("\nIs a valid binary tree ?")
	fmt.Println(isValidBinaryTree(root, math.MinInt, math.MaxInt))

}
