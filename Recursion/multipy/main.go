package main

import "fmt"

func multiplyArray(arr []int, start int, end int) int {
	if start > end {
		return 1
	}
	return arr[start] * multiplyArray(arr, start+1, end)
}

func main() {
	intArray := []int{2, 3, 5, 7, 11}

	fmt.Println("Original array:", intArray)

	result := multiplyArray(intArray, 0, len(intArray)-1)

	fmt.Println("Product of all elements:", result)
}
