package main

import "fmt"

type NodeDLL struct {
	data int
	next *NodeDLL
	prev *NodeDLL
}

type DoubleLL struct {
	head *NodeDLL
	tail *NodeDLL
}

func NewDoubleLL() *DoubleLL {
	return &DoubleLL{}
}

func main() {
	dll := NewDoubleLL()
	dll.AddingIntoDLL(1)
	dll.AddingIntoDLL(2)
	dll.AddingIntoDLL(3)
	dll.AddingIntoDLL(4)
	dll.DisplayDataForward()
	dll.DisplayDataReverse()
}

func (dll *DoubleLL) AddingIntoDLL(data int) {
	newNode := NodeDLL{data: data, next: nil, prev: nil}
	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
		return
	} else {
		newNode.prev = dll.tail
		dll.tail.next = &newNode
		dll.tail = &newNode
	}
}

func (dll *DoubleLL) DisplayDataForward() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (dll *DoubleLL) DisplayDataReverse() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
	fmt.Println()
}
