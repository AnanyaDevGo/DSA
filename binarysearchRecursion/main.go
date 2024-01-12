package main

import (
	"fmt"
)

func BinarySearch(arr []int, target, low, high int) int {
	if low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] > target {
			return BinarySearch(arr, target, low, mid-1)
		}

		return BinarySearch(arr, target, mid+1, high)
	}

	return -1
}

func main() {
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 6

	index := BinarySearch(sortedArray, target, 0, len(sortedArray)-1)

	if index != -1 {
		fmt.Printf("Target %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array.\n", target)
	}
}