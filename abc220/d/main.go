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
	As := make([]int, N+1)
	dp := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		As[i] = nextInt()
		dp[i] = make([]int, 10)
	}
	dp[0] = make([]int, 10)
	dp[1][As[1]] = 1

	for i := 1; i <= N-1; i++ {
		for j := 0; j <= 9; j++ {
			dp[i+1][(j+As[i+1])%10] += dp[i][j] % 998244353
			dp[i+1][(j+As[i+1])%10] %= 998244353
			dp[i+1][(j*As[i+1])%10] += dp[i][j] % 998244353
			dp[i+1][(j*As[i+1])%10] %= 998244353
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(dp[N][i])
	}
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
