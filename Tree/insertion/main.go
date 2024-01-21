package main

import (
	"fmt"
	// "math"
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

// func findMin(root *Node) *Node {
// 	current := root
// 	for current.Left != nil {
// 		current = current.Left
// 	}
// 	return current
// }

// func findMax(root *Node) *Node {
// 	current := root
// 	for current.Right != nil {
// 		current = current.Right
// 	}
// 	return current
// }

// func preorderTraversal(root *Node) {
// 	if root != nil {
// 		fmt.Printf("%d ", root.Key)
// 		preorderTraversal(root.Left)
// 		preorderTraversal(root.Right)
// 	}
// }

// func inorderTraversal(root *Node) {
// 	if root != nil {
// 		inorderTraversal(root.Left)
// 		fmt.Printf("%d ", root.Key)
// 		inorderTraversal(root.Right)
// 	}
// }

// func postorderTraversal(root *Node) {
// 	if root != nil {
// 		postorderTraversal(root.Left)
// 		postorderTraversal(root.Right)
// 		fmt.Printf("%d ", root.Key)
// 	}
// }

// func removeNode(root *Node, key int) *Node {
// 	if root == nil {
// 		return root
// 	}

// 	if key < root.Key {
// 		root.Left = removeNode(root.Left, key)
// 	} else if key > root.Key {
// 		root.Right = removeNode(root.Right, key)
// 	} else {
// 		if root.Left == nil {
// 			return root.Right
// 		} else if root.Right == nil {
// 			return root.Left
// 		}
// 		root.Key = findMin(root.Right).Key
// 		root.Right = removeNode(root.Right, root.Key)
// 	}

// 	return root
// }

// func closest(root *Node, target int) int {
// 	closestVal := math.Inf(1)
// 	for root != nil {
// 		if diff := float64(root.Key - target); math.Abs(diff) < math.Abs(closestVal-float64(target)) {
// 			closestVal = float64(root.Key)
// 		}
// 		if target < root.Key {
// 			root = root.Left
// 		} else {
// 			root = root.Right
// 		}
// 	}
// 	return int(closestVal)
// }

// func isValidBinaryTree(root *Node, minVal, maxVal int) bool {
// 	if root == nil {
// 		return true
// 	}
// 	if !(minVal < root.Key && root.Key < maxVal) {
// 		return false
// 	}
// 	return isValidBinaryTree(root.Left, minVal, root.Key) && isValidBinaryTree(root.Right, root.Key, maxVal)
// }

// func bfs(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	queue := []*Node{root}
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		fmt.Printf("%d ", current.Key)

// 		if current.Left != nil {
// 			queue = append(queue, current.Left)
// 		}
// 		if current.Right != nil {
// 			queue = append(queue, current.Right)
// 		}
// 	}
// }

func main() {
	var root *Node
	keys := []int{15, 10, 20, 8, 12, 17, 25}

	for _, key := range keys {
		root = insert(root, key)
	}

	// fmt.Println("In-order traversal:")
	// inorderTraversal(root)
	fmt.Println("\nSearch for 12:")
	result := search(root, 12)
	if result != nil {
		fmt.Println(result.Key)
	} else {
		fmt.Println("Not found")
	}

	// fmt.Println("\nMinimum value:")
	// minNode := findMin(root)
	// if minNode != nil {
	// 	fmt.Println(minNode.Key)
	// } else {
	// 	fmt.Println("Tree is empty")
	// }

	// fmt.Println("\nMaximum value:")
	// maxNode := findMax(root)
	// if maxNode != nil {
	// 	fmt.Println(maxNode.Key)
	// } else {
	// 	fmt.Println("Tree is empty")
	// }

	// fmt.Println("\nPre-order traversal:")
	// preorderTraversal(root)

	// fmt.Println("\nPost-order traversal:")
	// postorderTraversal(root)

	// fmt.Println("\nRemoving node with key 15:")
	// root = removeNode(root, 15)
	// inorderTraversal(root)

	// fmt.Println("\nClosest value to 13:")
	// fmt.Println(closest(root, 13))

	// fmt.Println("\nIs a valid binary tree?")
	// fmt.Println(isValidBinaryTree(root, math.MinInt, math.MaxInt))

	// fmt.Println("\nBFS traversal:")
	// bfs(root)
}
