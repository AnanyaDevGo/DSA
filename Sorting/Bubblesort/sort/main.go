package main

import "fmt"

func main() {
	data := []int{64, 25, 12, 22, 11}

	fmt.Println("Unsorted Array:", data)

	bubbleSort(data)

	fmt.Println("Sorted Array:", data)
}

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j <  n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
