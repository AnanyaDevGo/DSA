package main

import "fmt"

func replaceAlphabet(str string, n int) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		ascii := int(char)
		if ('A' <= char && char <= 'Z' && ascii <= 90-n) || ('a' <= char && char <= 'z' && ascii <= 122-n) {
			ascii += n
			char = rune(ascii)
		}
		runes[i] = char
	}
	return string(runes)
}
func main() {
	input := "Hello, World!"
	n := 2
	result := replaceAlphabet(input, n)
	fmt.Println(result)
}
