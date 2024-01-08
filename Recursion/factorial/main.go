package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	number := 5
	result := factorial(number)

	fmt.Printf("Factorial of %d is: %d\n", number, result)
}
