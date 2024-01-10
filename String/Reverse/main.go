package main

import "fmt"

func Reverse(str string) string {
	char := []rune(str)
	for i, j := 0, len(char)-1; i < j; i, j = i+1, j-1 {
		char[i], char[j] = char[j], char[i]
	}
	return string(char)

}

func main() {
	c := Reverse("riyas")
	fmt.Println("reversed", c)
}
