package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	data := []int{64, 25, 12, 22, 11}

	fmt.Println("Unsorted Array:", data)

	selectionSort(data)

	fmt.Println("Sorted Array:", data)
}
