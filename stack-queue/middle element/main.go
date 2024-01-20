package main

import (
	"container/list"
	"fmt"
)


type Stack struct {
items []int
}

func findMiddle(stack *list.List) interface{}{
if stack.Len() == 0 {
return nil
}


slow := stack.Front()
fast := stack.Front()

for fast != nil && fast.Next() != nil {
slow = slow.Next()
fast = fast.Next().Next()
}
return slow.Value
}

func main (){
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	stack.PushBack(4)
	stack.PushBack(5)
	
	middle := findMiddle(stack)
	 if middle != nil {
	 fmt.Println("Middle element :", middle)
	 } else {
	 fmt.Println("The stack is empty")
	 
	 }
} 
