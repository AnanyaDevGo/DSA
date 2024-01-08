package main

import "fmt"

func arraySum(arr []int, index int) int {
	if index == len(arr) {
		return 0
	}

	return arr[index] + arraySum(arr, index+1)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := arraySum(numbers, 0)
	fmt.Printf("The sum of the array %v is %d\n", numbers, result)
}
