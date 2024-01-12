package main

import "fmt"

func linearSearch(arr []int, target int) int {
    for i, value := range arr {
        if value == target {
            return i 
        }
    }
    return -1 
}

func main() {
    array := []int{5, 2, 9, 1, 5, 6}
    target := 1

    index := linearSearch(array, target)

    if index != -1 {
        fmt.Printf("Linear Search: %d found at index %d\n", target, index)
    } else {
        fmt.Printf("Linear Search: %d not found in the array\n", target)
    }
}
