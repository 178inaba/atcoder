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

	if total%10 != 0 {
		fmt.Println(total)
		return
	}

	minS := 101
	for _, s := range ss {
		if s%10 != 0 {
			minS = min(minS, s)
		}
	}

	if minS == 101 {
		fmt.Println(0)
		return
	}
	fmt.Println(total - minS)
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
