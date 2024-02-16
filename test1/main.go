package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func inOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	fmt.Println(root.Data, "")
	inOrderTraversal(root.Right)
}
func preOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data, " ")
	preOrderTraversal(root.Left)
	preOrderTraversal(root.Right)
}
func main() {
	root := &Node{Data: 1}
	root.Left = &Node{Data: 2}
	root.Right = &Node{Data: 3}
	root.Left.Left = &Node{Data: 4}
	root.Left.Right = &Node{Data: 5}
	root.Right.Left = &Node{Data: 6}
	fmt.Println("In-Order traversal:")
	inOrderTraversal(root)
	fmt.Println()
	fmt.Println("Pre-Order Traversal:")
	preOrderTraversal(root)
	fmt.Println()
}
