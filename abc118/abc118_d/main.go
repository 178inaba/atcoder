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
	var As []int
	for i := 0; M > i; i++ {
		As = append(As, nextInt())
	}

	numMap := map[int]int{1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6}
	priorities := []int{1, 7, 4, 5, 3, 2, 9, 6, 8}
	var currentPriority []int
	for _, priority := range priorities {
		for _, A := range As {
			if A == priority {
				currentPriority = append(currentPriority, A)
			}
		}
	}

	parts := map[int]int{}
	for N > 0 {
		fmt.Println(N)
		for _, cp := range currentPriority {
			if N > numMap[cp] {
				N -= numMap[cp]
				parts[cp]++
				break
			}
		}
	}

	fmt.Println(parts)
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
