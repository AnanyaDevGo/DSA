package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var less, equal, greater []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	return append(append(quicksort(less), equal...), quicksort(greater)...)
}

func main() {
	data := []int{64, 25, 12, 22, 11}

	fmt.Println("Unsorted Array:", data)

	sortedData := quicksort(data)

	fmt.Println("Sorted Array:", sortedData)
}
