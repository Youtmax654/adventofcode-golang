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

	santaPosition := coordinate{0, 0}
	robotPosition := coordinate{0, 0}
	houses := map[coordinate]bool{santaPosition: true}

	for i, direction := range string(directions) {
		if i % 2 == 0 {
			// p := santaPosition
			// switch direction {
			// case '^':
			// 	p.y++
			// case 'v':
			// 	p.y--
			// case '>':
			// 	p.x++
			// case '<':
			// 	p.x--
			// }
			// santaPosition = p
			// houses[p] = true
			// Corrected by ChatGPT
			santaPosition = move(santaPosition, direction)
			houses[santaPosition] = true
		} else {
			// p := robotPosition
			// switch direction {
			// case '^':
			// 	p.y++
			// case 'v':
			// 	p.y--
			// case '>':
			// 	p.x++
			// case '<':
			// 	p.x--
			// }
			// robotPosition = p
			// houses[p] = true
			// Corrected by ChatGPT
			robotPosition = move(robotPosition, direction)
			houses[robotPosition] = true
		}
	}

	fmt.Printf("There is %d houses that received at least one present.\n", len(houses))
}

// Added by ChatGPT
func move(pos coordinate, direction rune) coordinate {
	switch direction {
	case '^':
		pos.y++
	case 'v':
		pos.y--
	case '>':
		pos.x++
	case '<':
		pos.x--
	}
	return pos
}