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
	K := nextInt()
	var isOdd bool
	if K%2 == 0 {
		isOdd = true
	}

	var cnt int
	for i := 1; i < K; i *= 2 {
		cnt++
	}

	var kakuritsu float64
	kaisu := N
	if N > K {
		if !isOdd {
			kaisu = cnt + 1
		} else {
			kaisu = cnt
		}

		kakuritsu = float64(N-kaisu) / float64(N)
	}

	for i := 0; i < cnt; i++ {
		kakuritsu += (float64(1) / float64(N)) * math.Pow(0.5, float64(cnt-i))
	}
	if !isOdd {
		kakuritsu += (float64(1) / float64(N)) * math.Pow(0.5, 1)
	}

	fmt.Printf("%.10f\n", kakuritsu)
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
