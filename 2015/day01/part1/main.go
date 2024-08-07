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

	for _, instruction := range instructions {
		if instruction == "(" {
			floor += 1
			continue
		}

		if instruction == ")" {
			floor -= 1
			continue
		}
	}

	fmt.Println("The instructions take Santa to the floor", floor)
}