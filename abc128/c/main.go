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
	ss := make([][]int, M)
	for i := 0; i < M; i++ {
		k := nextInt()
		ss[i] = make([]int, k)
		for j := 0; j < k; j++ {
			ss[i][j] = nextInt()
		}
	}

	ps := make([]int, M)
	for i := 0; i < M; i++ {
		ps[i] = nextInt()
	}

	allSw := map[int]bool{}
	var ans int
	for i := 0; ; i++ {
		if i != 0 {
			inc(allSw, 1)
		}

		isOK := true
		for j := 0; j < M; j++ {
			var onCnt int
			for _, s := range ss[j] {
				if allSw[s] {
					onCnt++
				}
			}
			if onCnt%2 != ps[j] {
				isOK = false
				break
			}
		}
		if isOK {
			ans++
		}

		isEnd := true
		for k := 1; k <= N; k++ {
			if !allSw[k] {
				isEnd = false
			}
		}
		if isEnd {
			break
		}
	}

	fmt.Println(ans)
}

func inc(m map[int]bool, i int) {
	if m[i] == true {
		m[i] = false
		inc(m, i+1)
		return
	}

	m[i] = true
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
