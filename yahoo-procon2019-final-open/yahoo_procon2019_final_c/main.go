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
	squares := map[int]map[int]bool{}
	for i := 0; N > i; i++ {
		H := nextInt()
		W := nextInt()
		R := nextInt()
		C := nextInt()
		for p := 1; H >= p; p++ {
			currentC := C
			for q := 1; W >= q; q++ {
				if squares[R] == nil {
					squares[R] = map[int]bool{}
				}

				if (p+q)%2 == 0 {
					if squares[R][currentC] {
						squares[R][currentC] = false
					} else {
						squares[R][currentC] = true
					}
				}

				currentC++
			}

			R++
		}
	}

	var result int
	for _, row := range squares {
		for _, square := range row {
			if square {
				result++
			}
		}
	}

	fmt.Println(result)
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
