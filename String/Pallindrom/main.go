package main

import "fmt"

func Palindrome(str string) {
	c := 0
	for i, j := 0, len(str)-1; i <= len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			c++
		}
	}
	if c == 0 {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
func main() {
	Palindrome("MALAYALAM")
	Palindromee("Ananya")
}
func Palindromee(str string) {
	c := 0
	char := []rune(str)
	for i, j := 0, len(str)-1; i <= len(str)/2; i, j = i+1, j-1 {
		if char[i] != char[j] && char[i] != char[j]-32 && char[i] != char[j]+32 {
			c++
		}
	}
	if c == 0 {
		fmt.Println("palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
