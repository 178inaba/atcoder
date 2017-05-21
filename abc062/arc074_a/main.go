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
		b1 := (y / 2) * (x - i)
		b2 := (x - i) / 2 * y
		c1 := x*y - a - b1
		c2 := x*y - a - b2
		ans = min(ans, max(a, max(b1, c1))-min(a, min(b1, c1)))
		ans = min(ans, max(a, max(b2, c2))-min(a, min(b2, c2)))
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
