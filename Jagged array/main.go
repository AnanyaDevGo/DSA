package main

import "fmt"

func main() {

	jaggedArray := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
		{10},
	}

	fmt.Println(jaggedArray[0]) // [1 2 3]
	fmt.Println(jaggedArray[1]) // [4 5]
	fmt.Println(jaggedArray[2]) // [6 7 8 9]
	fmt.Println(jaggedArray[3]) // [10]
}
