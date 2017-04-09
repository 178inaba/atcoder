package main

import "math"

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func diff(x, y int) int {
	return abs(-(x) + y)
}

func countRune(s string, r rune) int {
	var cnt int
	for _, sr := range s {
		if sr == r {
			cnt++
		}
	}

	return cnt
}
