package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := nextInt()
	var cnt, hasACnt, hasBCnt int
	for i := 0; i < N; i++ {
		S := nextString()
		cnt += strings.Count(S, "AB")

		if hasB := strings.HasPrefix(S, "B"); hasB {
			hasBCnt++
		}

		if hasA := strings.HasSuffix(S, "A"); hasA {
			hasACnt++
		}
	}

	m := min(hasACnt, hasBCnt)
	if hasACnt == hasBCnt {
		m--
	}

	fmt.Println(cnt + m)
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
