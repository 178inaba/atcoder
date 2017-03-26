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

	an := make([]int, 0, N)
	bn := make([]int, 0, N)
	for i := 0; i < N; i++ {
		an = append(an, nextInt())
		bn = append(bn, nextInt())
	}

	cm := make([]int, 0, M)
	dm := make([]int, 0, M)
	for i := 0; i < M; i++ {
		cm = append(cm, nextInt())
		dm = append(dm, nextInt())
	}

	keys := make([]int, 0, N)
	for i := 0; i < N; i++ {
		result, k := 0, 0
		for j := 0; j < M; j++ {
			r := abs(an[i]-cm[j]) + abs(bn[i]-dm[j])
			if k == 0 || r < result {
				result = r
				k = j + 1
			}
		}

		keys = append(keys, k)
	}

	for _, k := range keys {
		fmt.Println(k)
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
	if x < y {
		return x
	}

	return y
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.
