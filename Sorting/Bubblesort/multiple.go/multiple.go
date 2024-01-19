package main

import "fmt"

func bubbleSort(arr []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSort(arr, n-1)
}

func multiplyArray(arr []int, n int) int {
	if n == 0 {
		return 1
	}
	return arr[n-1] * multiplyArray(arr, n-1)
}

func main() {
	intArray := []int{5, 3, 7, 2, 8}

	fmt.Println("Original array:", intArray)

	bubbleSort(intArray, len(intArray))

	fmt.Println("Sorted array:", intArray)

	result := multiplyArray(intArray, len(intArray))
	fmt.Println("Product of all elements:", result)
}
