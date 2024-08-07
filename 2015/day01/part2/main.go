package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	instructions := strings.Split(string(input), "")
	floor := 0
	position := 0

	for i, instruction := range instructions {
		if floor == -1 {
			position = i
			break
		}

		if instruction == "(" {
			floor += 1
			continue
		}

		if instruction == ")" {
			floor -= 1
			continue
		}
	}

	fmt.Println("The position of the character that causes Santa to first enter the basement is position", position)
}