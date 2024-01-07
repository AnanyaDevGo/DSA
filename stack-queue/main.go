package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// push = add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// pop = remove a value at end and return it
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
