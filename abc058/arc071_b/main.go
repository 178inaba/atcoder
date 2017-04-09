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
	m := nextInt()
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt()
	}
	ys := make([]int, m)
	for i := 0; i < m; i++ {
		ys[i] = nextInt()
	}
	p := int(math.Pow10(9)) + 7

	var X, Y int
	for i, x := range xs {
		X += i*x - (n-(i+1))*x
		X %= p
	}
	for i, y := range ys {
		Y += i*y - (m-(i+1))*y
		Y %= p
	}

	fmt.Println(X * Y % p)
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

func diff(x, y int) int {
	return abs(-(x) + y)
}

// ---------- Util.
