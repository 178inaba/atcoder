package main

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
