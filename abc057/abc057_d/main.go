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

	// Make dp.
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	// Calc N and k.
	An, Ak := 0, 0
	for i := 0; i < N; i++ {
		if vs[i] == vs[A-1] {
			An++
			if i < A {
				Ak++
			}
		}
	}

	cnt := 0
	if Ak == A {
		for ; Ak <= B; Ak++ {
			if Ak >= len(dp[An]) {
				break
			}

			cnt += dp[An][Ak]
		}
	} else {
		cnt = dp[An][Ak]
	}

	fmt.Printf("%f\n%d\n", avg, cnt)
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
