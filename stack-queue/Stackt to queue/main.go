package main

import "fmt"

func StackToQueue(stack []int) map[int]int {
	queue := make(map[int]int)
	for i, v := range stack {
		queue[len(stack)-i-1] = v
	}
	return queue
}

func main() {
	stack := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("the stack", stack)
	queue := StackToQueue(stack)
	fmt.Println("the quene", queue)

}
