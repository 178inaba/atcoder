package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	n := nextInt()
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = nextString()
	}

	var ans []rune
	for r := 'a'; r <= 'z'; r++ {
		small := 50
		for _, s := range S {
			small = min(small, countRune(s, r))
		}

		for i := 0; i < small; i++ {
			ans = append(ans, r)
		}
	}

	fmt.Println(string(ans))
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

// Input. ----------

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func nextString() string {
	sc.Scan()
	if err := sc.Err(); err != nil {
		panic(err)
	}

	return sc.Text()
}

// ---------- Input.

// Util. ----------

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.
