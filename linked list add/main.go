package main

import "fmt"

type Node struct{
	data  int
	next *Node
}

type LinkedList struct{
	head *Node
}
func main (){
	list := LinkedList{}
	list.Appendlist(1)
	list.Appendlist(2)
	list.Appendlist(3)
	list.Appendlist(4)
	list.Displaylist()
}

func (list *LinkedList) Appendlist(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func(list *LinkedList) Displaylist(){
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func (li *LinkedList) InsertAfter(previousNode *Node, data int){
	if previousNode == nil {
		fmt.Println("previous node cannot be nil")
	}

	newNode := &Node{data: data, next: previousNode.next}
	previousNode.next = newNode

}