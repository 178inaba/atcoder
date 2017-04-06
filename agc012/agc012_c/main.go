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
	var n int
	var p []int
	for n = 40; n > 0; n-- {
		if N >= (1<<uint(n))-1 {
			break
		}
	}
	for i := 1; i <= n; i++ {
		p = append(p, i)
	}
	N -= (1 << uint(n)) - 1

	for i := n - 1; i >= 0; i-- {
		if N < 1<<uint(i) {
			continue
		}
		n++
		p = append(p[:i], append([]int{n}, p[i:]...)...)
		N -= 1 << uint(i)
	}

	fmt.Println(len(p) * 2)
	for i := 0; i < len(p); i++ {
		fmt.Printf("%d ", i+1)
	}

	for _, v := range p {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
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
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.
