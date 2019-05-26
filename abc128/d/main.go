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
	N, K := nextInt(), nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}

	var ans int
	for l := 0; l <= K; l++ {
		for r := 0; r <= K-l; r++ {
			if l+r > N {
				continue
			}

			d := K - l - r
			var now int
			var s []int
			for i := 0; i < l; i++ {
				now += a[i]
				s = append(s, a[i])
			}
			for i := N - r; i < N; i++ {
				now += a[i]
				s = append(s, a[i])
			}

			sort.Ints(s)
			for i := 0; i < d; i++ {
				if i >= len(s) {
					break
				} else if s[i] > 0 {
					break
				}

				now -= s[i]
			}

			ans = max(ans, now)
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
