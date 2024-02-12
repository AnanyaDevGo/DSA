// package main

// import "fmt"

// func countDuplicates(arr1, arr2 []int) int {
//     set := make(map[int]bool)
//     duplicates := 0
//     for _, num := range arr1 {
//         set[num] = true
//     }
//     for _, num := range arr2 {
//         if set[num] {
//             duplicates++
//         }
//     }
//     return duplicates
// }

//	func main() {
//	    arr1 := []int{1, 2, 5, 3, 6, 4}
//	    arr2 := []int{4, 6, 9, 8, 4, 3}
//	    result := countDuplicates(arr1, arr2)
//	    fmt.Println("Number of duplicated elements:", result)
//	}
// package main

// import "fmt"

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

//	func main() {
//		str := "hello"
//		reversed := reverseString(str)
//		fmt.Println("Reversed string:", reversed)
//	}
package main

import "fmt"

func countVowels(s string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for _, char := range s {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	str := "hello"
	vowelsCount := countVowels(str)
	fmt.Println("Number of vowels:", vowelsCount)
}
