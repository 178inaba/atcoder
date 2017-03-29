package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := nextInt()
	A := nextInt()
	B := nextInt()

	vs := make([]int, N)
	for i := 0; i < N; i++ {
		vs[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(vs)))

	var avg float64
	for i := 0; i < A; i++ {
		avg += float64(vs[i])
	}

	avg /= float64(A)

	fmt.Printf("%f\n", avg)

	// Debug.
	B++
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
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.
