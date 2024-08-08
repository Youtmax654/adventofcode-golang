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
		lwh := strings.Split(dimension, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		// lw := l + l + w + w
		// wh := w + w + h + h
		// hl := h + h + l + l
		// Corrected by ChatGPT
		lw := 2 * (l + w)
		wh := 2 * (w + h)
		hl := 2 * (h + l)

		shortest := min(lw, wh, hl)
		bow := l * w * h

		total += shortest + bow
	}

	fmt.Println("The total feet of ribbon they should order is", total)
}