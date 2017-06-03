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
	for i := 0; i < N; i++ {
		ss[i] = nextInt()
	}

	var ans int
	for i := 0; i < N; i++ {
		if ss[i]%10 != 0 {
			ans = max(ans, ss[i])
		}
		total := ss[i]
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			total += ss[j]
			if total%10 != 0 {
				ans = max(ans, total)
			}
			if (ss[i]+ss[j])%10 != 0 {
				ans = max(ans, ss[i]+ss[j])
			}
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
