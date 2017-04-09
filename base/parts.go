package main

func countRune(s string, r rune) int {
	var cnt int
	for _, sr := range s {
		if sr == r {
			cnt++
		}
	}

	return cnt
}
