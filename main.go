package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Println("Before Remove:", arr)
	newLength := removeDuplicates(arr)
	fmt.Println("After Remove:", arr[:newLength])
}

func removeDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var new []int
	new = append(new, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			new = append(new, arr[i])
		}
	}
	copy(arr, new)
	length := len(new)
	return length

}
