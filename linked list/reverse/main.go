package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev, next *ListNode
    current := head

    for current != nil {
        next = current.Next
        current.Next = prev 
        
        prev = current
        current = next
    }

    return prev
}

func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    node4 := &ListNode{Val: 4, Next: nil}
    node3 := &ListNode{Val: 3, Next: node4}
    node2 := &ListNode{Val: 2, Next: node3}
    node1 := &ListNode{Val: 1, Next: node2}

    fmt.Println("Original linked list:")
    printList(node1)

    reversedList := reverseList(node1)

    fmt.Println("Reversed linked list:")
    printList(reversedList)
}

