package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	// hash := md5.Sum([]byte(input))
	// hexHash := hex.EncodeToString(hash[:])
	// Removed by ChatGPT

	i := 1
	// for !strings.HasPrefix(hexHash, "000000") {
	// Corrected by ChatGPT
	for {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hexHash := hex.EncodeToString(hash[:])
		fmt.Println(i)

		// Added by ChatGPT
		if strings.HasPrefix(hexHash, "000000") {
			break
		}

		i++
	}

	// fmt.Printf("The answer is %d\n.", i - 1)
	// Corrected by ChatGPT
	fmt.Printf("The answer is %d\n.", i)
}