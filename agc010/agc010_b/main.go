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
	if N == 1 {
		fmt.Println("YES")
		return
	}

	A := make([]int, 0, N)
	var total int
	for i := 0; i < N; i++ {
		v := nextInt()
		A = append(A, v)
		total += v
	}

	removeNum := N * (N + 1) / 2
	if total%removeNum != 0 {
		fmt.Println("NO")
		return
	}

	total /= removeNum
	for i := 0; i < N; i++ {
		t := A[i] - A[(i+1)%N] + total
		if t < 0 || t%N != 0 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
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

// ---------- Util.
