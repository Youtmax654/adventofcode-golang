package main

import (
	"fmt"
	"os"
)

type coordinate struct {
	x int
	y int
}

func main() {
	directions, _ := os.ReadFile("input.txt")

	current := coordinate{0, 0}

	// visited := make(map[coordinate]bool)
	// visited[current] = true
	// Corrected by ChatGPT
	visited := map[coordinate]bool{current: true}

	for _, direction := range string(directions) {
		switch direction {
		case '^':
			current.y++
		case 'v':
			current.y--
		case '>':
			current.x++
		case '<':
			current.x--
		}

		// _, ok := visited[current]
		// if !ok {
		// 	visited[current] = true
		// }
		// Corrected by ChatGPT
		visited[current] = true
	}

	// fmt.Println("There is", len(visited), "houses that receive at least one present.")
	fmt.Printf("There is %d houses that received at least one present.\n", len(visited))
}