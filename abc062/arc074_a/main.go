package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	W := nextInt()

	s1 := calc(H, W)
	s2 := calc(W, H)
	fmt.Println(min(s1, s2))
}

func calc(x, y int) int {
	ans := math.MaxInt64
	for i := 1; i < x; i++ {
		a := y * i
		b := (y / 2) * (x - i)
		c := x*y - a - b
		ans = min(ans, max(a, max(b, c))-min(a, min(b, c)))
	}

	for i := 1; i < x; i++ {
		a := y * i
		b := (x - i) / 2 * y
		c := x*y - a - b
		ans = min(ans, max(a, max(b, c))-min(a, min(b, c)))
	}

	return ans
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

// ---------- Util.
