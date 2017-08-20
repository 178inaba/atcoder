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
	N := nextInt()
	as := make([]int, N)
	for i := 0; i < N; i++ {
		as[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(as)))

	var before int
	var ans []int
	for _, a := range as {
		if a == before {
			ans = append(ans, a)
			if len(ans) == 2 {
				break
			}
			before = 0
			continue
		}
		before = a
	}
	if len(ans) == 2 {
		fmt.Println(ans[0] * ans[1])
		return
	}
	fmt.Println(0)
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
