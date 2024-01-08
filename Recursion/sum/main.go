package main

import "fmt"

func sum(n int) int {
    if n == 0 {
        return 0
    }

    return n + sum(n-1)
}

func main() {
    n := 5
    result := sum(n)
    fmt.Printf("The sum of numbers from 1 to %d is %d\n", n, result)
}
