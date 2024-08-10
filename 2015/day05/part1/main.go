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
		if !hasThreeVowels(s) {
			continue
		}

		if !hasRepeatedLetters(s) {
			continue
		}

		if hasDisallowedSubstrings(s) {
			continue
		}

		valid++
	}

	fmt.Printf("There is %d strings that are nice.\n", valid)
}

// Corrected by ChatGPT
func hasThreeVowels(s string) bool {
	// vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowels := "aeiou"
	count := 0
	for _, l := range s {
		// for _, vowel := range vowels {
		// 	if l == vowel {
		// 		count++
		// 	}
		// }
		if strings.ContainsRune(vowels, l) {
			count++
		}
		if count >= 3 {
			return true
		}
	}

	// return count >= 3
	return false
}

func hasRepeatedLetters(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

// Corrected by ChatGPT
func hasDisallowedSubstrings(s string) bool {
	substrings := []string{"ab", "cd", "pq", "xy"}
	for i := 0; i < len(s) - 1; i++ {
		for _, substring := range substrings {
			// if string(s[i]) + string(s[i+1]) == substring {
			// 	return true
			// }
			if strings.Contains(s, substring) {
				return true
			}
		}
	}

	return false
}