package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	dimensions := strings.Split(string(input), "\n")

	total := 0

	for _, dimension := range dimensions {
		smallest := 0
		lwh := strings.Split(dimension, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		lw := l * w
		wh := w * h
		hl := h * l

		// smallest = lw
		// if wh < smallest {
		// 	smallest = wh
		// }
		// if hl < smallest {
		// 	smallest = hl
		// }
		// Corrected by ChatGPT
		smallest = min(lw, wh, hl)

		s := 2*lw + 2*wh + 2*hl
		total += smallest + s
	}

	fmt.Println("The total square feet of wrapping paper they should order is", total)
}