package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var matchNum = []int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func main() {
	N := nextInt()
	M := nextInt()
	var As []int
	for i := 0; M > i; i++ {
		As = append(As, nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(As)))

	dp := make([]int, N+1)
	dp[0] = 0
	for i := 1; N >= i; i++ {
		maxDigit := int(math.Inf(-1))
		for _, A := range As {
			index := i - matchNum[A]
			if index < 0 || dp[index] < 0 {
				continue
			}

			maxDigit = max(maxDigit, dp[index]+1)
		}

		dp[i] = maxDigit
	}

	remain := dp[N]
	match := N
	for match > 0 {
		for _, A := range As {
			if match-matchNum[A] < 0 {
				continue
			}

			if dp[match-matchNum[A]] == remain-1 {
				fmt.Print(A)
				match -= matchNum[A]
				remain--
				break
			}
		}
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
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.
