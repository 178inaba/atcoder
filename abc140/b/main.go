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
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = nextInt()
	}

	bs := make([]int, n)
	for i := 0; i < n; i++ {
		bs[i] = nextInt()
	}

	cs := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		cs[i] = nextInt()
	}

	var score int
	before := -1
	for _, a := range as {
		score += bs[a-1]
		if before+1 == a {
			score += cs[before-1]
		}

		before = a
	}

	fmt.Println(score)
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

func nextFloat() float64 {
	sc.Scan()
	f64, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}

	return f64
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

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.
