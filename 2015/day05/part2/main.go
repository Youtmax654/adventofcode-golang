package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	array := strings.Split(string(input), "\n")

	valid := 0
	for _, s := range array {
		if !hasTwoPair(s) {
			continue
		}

		if !repeatWithLetterBetween(s) {
			continue
		}

		valid++
	}

	fmt.Printf("There is %d strings that are nice.\n", valid)
}

// func has_two_pair(s string) bool {
// 	for i := 0; i < len(s) - 1; i++  {
// 		pair := string(s[i]) + string(s[i+1])

// 		for j := i + 2; j < len(s) - 1; j++ {
// 			if pair == string(s[j]) + string(s[j+1]) {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }
// Optimised by ChatGPT
// To avoid using two for loop
func hasTwoPair(s string) bool {
	pairs := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]

		// Check if the pair has already been seen and isn't overlapping
		if j, exists := pairs[pair]; exists && i > j+1 {
			return true
		}

		// Store the index of the current pair
		pairs[pair] = i
	}

	return false
}

func repeatWithLetterBetween(s string) bool {
	for i := 0; i < len(s) - 2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}