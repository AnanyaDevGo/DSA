package main

import "fmt"

func binarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1

    for low <= high {
        mid := (low + high) / 2

        if arr[mid] == target {
            return mid // Return the index of the target if found
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1 // Return -1 if the target is not found
}

func main() {
    array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    target := 6

    index := binarySearch(array, target)

    if index != -1 {
        fmt.Printf("Binary Search: %d found at index %d\n", target, index)
    } else {
        fmt.Printf("Binary Search: %d not found in the array\n", target)
    }
}
