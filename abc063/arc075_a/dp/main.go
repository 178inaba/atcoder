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
	ss := make([]int, N)
	var total int
	for i := 0; i < N; i++ {
		ss[i] = nextInt()
		total += ss[i]
	}

	dp := make([][]bool, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]bool, total+1)
	}
	dp[0][0] = true
	for i := 0; i < N; i++ {
		for j := 0; j <= total; j++ {
			if dp[i][j] {
				dp[i+1][j] = true
				dp[i+1][j+ss[i]] = true
			}
		}
	}
	var ans int
	for score, ok := range dp[N] {
		if ok && score%10 != 0 {
			ans = max(ans, score)
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
