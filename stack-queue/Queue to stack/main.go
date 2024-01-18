package main

import "fmt"

func QueneToStack(queue []int) map[int]int {
	stack := make(map[int]int)
	for i, v := range queue {
		stack[i] = v
	}
	return stack
}
func main() {
	queue := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("the quene:", queue)
	stack := QueneToStack(queue)
	fmt.Println("the stack", stack)
}
