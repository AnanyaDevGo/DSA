package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	list := LinkedList{}
	list.Appendlist(1)
	list.Appendlist(2)
	list.Appendlist(3)
	list.Appendlist(4)
	list.Appendlist(5)
	list.Appendlist(6)
	list.Displaylist()
}

func (list *LinkedList) Appendlist(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (list *LinkedList) Displaylist() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next

	}
}
