package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N, A, B, C int
	l          []int
)

func main() {
	N = nextInt()
	A = nextInt()
	B = nextInt()
	C = nextInt()
	l = make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = nextInt()
	}

	fmt.Println(dfs(0, 0, 0, 0))
}

func dfs(cur, a, b, c int) int {
	if cur == N {
		if min(a, min(b, c)) > 0 {
			return abs(a-A) + abs(b-B) + abs(c-C) - 30
		}

		return int(math.MaxInt32)
	}

	ret0 := dfs(cur+1, a, b, c)
	ret1 := dfs(cur+1, a+l[cur], b, c) + 10
	ret2 := dfs(cur+1, a, b+l[cur], c) + 10
	ret3 := dfs(cur+1, a, b, c+l[cur]) + 10

	return min(min(ret0, ret1), min(ret2, ret3))
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
