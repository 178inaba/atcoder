package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N, M := nextInt(), nextInt()
	a := make([]int, N)
	for i := 0; i < M; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			s := nextInt()
			s--
			a[s] |= 1 << uint(i)
		}
	}

	var p int
	for i := 0; i < M; i++ {
		x := nextInt()
		p |= x << uint(i)
	}

	var ans int
	for s := 0; s < (1 << uint(N)); s++ {
		t := 0
		for i := 0; i < N; i++ {
			if s>>uint(i)&1 == 1 {
				t ^= a[i]
			}
		}

		if t == p {
			ans++
		}
	}

	fmt.Println(ans)
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
