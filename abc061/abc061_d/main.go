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
	M := nextInt()
	var a, b, c []int
	for l := 0; l < M; l++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, -nextInt())
	}

	dist := make([]int, N)
	for i := 1; i < N; i++ {
		dist[i] = math.MaxInt64
	}

	for l := 0; l < N-1; l++ {
		for i := 0; i < M; i++ {
			if dist[a[i]] == math.MaxInt64 {
				continue
			} else if dist[b[i]] > dist[a[i]]+c[i] {
				dist[b[i]] = dist[a[i]] + c[i]
			}
		}
	}

	ans := -dist[N-1]

	negative := make([]bool, N)
	for l := 0; l < N; l++ {
		for i := 0; i < M; i++ {
			if dist[a[i]] == math.MaxInt64 {
				continue
			} else if dist[b[i]] > dist[a[i]]+c[i] {
				dist[b[i]] = dist[a[i]] + c[i]
				negative[b[i]] = true
			} else if negative[a[i]] {
				negative[b[i]] = true
			}
		}
	}

	if negative[N-1] {
		fmt.Println("inf")
	} else {
		fmt.Println(ans)
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
