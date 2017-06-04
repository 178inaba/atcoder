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
	A := nextInt()
	B := nextInt()
	hs := make([]int, N)
	for i := 0; i < N; i++ {
		hs[i] = nextInt()
	}
	var under int
	upper := 1000000000
	for under+1 < upper {
		mid := (under + upper) / 2
		var cnt int
		for i := 0; i < N; i++ {
			if B*mid < hs[i] {
				cnt += int(math.Ceil(float64(hs[i]-B*mid) / float64(A-B)))
			}
		}
		if cnt <= mid {
			upper = mid
		} else {
			under = mid
		}
	}
	fmt.Println(upper)
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
