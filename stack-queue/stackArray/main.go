package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
}
func (s *Stack) Pop() int {
	n := len(s.data) - 1
	toRemove := s.data[n]
	s.data = s.data[:n]
	return toRemove
}
func (s *Stack) Display() {
	fmt.Println("Items are :", s.data)
}
func (s *Stack) Top() {
	topElement := s.data[len(s.data)-1]
	fmt.Println("Top Element is :", topElement)
}
func main() {
	Mystack := Stack{}
	Mystack.Push(20)
	Mystack.Push(40)
	Mystack.Push(50)
	Mystack.Push(60)
	Mystack.Push(70)
	Mystack.Push(80)
	Mystack.Push(90)
	Mystack.Display()
	a := Mystack.Pop()
	fmt.Println("Deleted Element is :", a)
	fmt.Println("After Pop Operation :", Mystack)
	Mystack.Top()
}
