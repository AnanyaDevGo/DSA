package main

import "fmt"

func mergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]
		mergeSort(left)
		mergeSort(right)

		i, j, k := 0, 0, 0

		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
			k++
		}

		for i < len(left) {
			arr[k] = left[i]
			i++
			k++
		}

		for j < len(right) {
			arr[k] = right[j]
			j++
			k++
		}
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	mergeSort(arr)
	fmt.Println("Sorted array:", arr)
}
