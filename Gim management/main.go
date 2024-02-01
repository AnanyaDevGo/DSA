package main

import "fmt"

type Node struct {
	ID    int
	Name  string
	Left  *Node
	Right *Node
}

func NewNode(id int, name string) *Node {
	return &Node{ID: id, Name: name, Left: nil, Right: nil}
}

func InorderTraversal(node *Node) {
	if node != nil {
		InorderTraversal(node.Left)
		fmt.Printf("ID:%d, Name: %s\n", node.ID, node.Name)
		InorderTraversal((node.Right))
	}
}

func main() {
	root := NewNode(1, "Anu")
	root.Left = NewNode(2, "Bibi")
	root.Right = NewNode(3, "Charls") 
	root.Left.Left = NewNode(4, "Devu")
	root.Left.Right = NewNode(5, "Edwin")

	fmt.Println("Result(ID and Name) ")
	InorderTraversal(root)
}
