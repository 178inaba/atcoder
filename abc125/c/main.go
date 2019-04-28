package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := nextInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	L := make([]int, N)
	R := make([]int, N)
	L[0] = A[0]
	R[N-1] = A[N-1]
	for i := 0; i < N-1; i++ {
		L[i+1] = gcd(L[i], A[i+1])
		R[N-1-i-1] = gcd(R[N-1-i], A[N-1-i-1])
	}

	var max int
	for i := 0; i < N; i++ {
		var g int
		if i == 0 {
			g = R[i+1]
		} else if i == N-1 {
			g = L[i-1]
		} else {
			g = gcd(L[i-1], R[i+1])
		}

		if g > max {
			max = g
		}
	}

	fmt.Println(max)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
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
