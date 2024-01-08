package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    number := 6
    result := fibonacci(number)

    fmt.Printf("Fibonacci sequence at position %d is: %d\n", number, result)
}
