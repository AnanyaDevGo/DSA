package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// Print a Node with value specified
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

// Delete a Node with value specified
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previeoudToDelete := l.head
	for previeoudToDelete.next.data != value {
		if previeoudToDelete.next.next == nil {
			return
		}
		previeoudToDelete = previeoudToDelete.next
	}
	previeoudToDelete.next = previeoudToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 20}
	node2 := &node{data: 22}
	node3 := &node{data: 30}
	node4 := &node{data: 11}
	node5 := &node{data: 10}
	node6 := &node{data: 39}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(10)
	mylist.printListData()
	// mylist.deleteWithValue(39)
	// mylist.printListData()
	// mylist.deleteWithValue(103)
	// mylist.printListData()
	// emptylist := linkedList{}
	// emptylist.deleteWithValue(10)
	// mylist.printListData()
}
