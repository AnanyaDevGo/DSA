package main

import "fmt"

func main() {
	arr := []int{2, 5, 8, 45, 45, 45, 2, 2, 6, 6, 7}
	fmt.Println("Before Delete Array is :", arr)
	Result := removeDuplicates(arr)
	fmt.Println("After Delete Array is :", Result)
}
func removeDuplicates(arr []int) []int {
	table := make(map[int]bool)
	var Array []int
	for _, v := range arr {
		if !table[v] {
			Array = append(Array, v)
			table[v] = true
		}
	}
	return Array
}
