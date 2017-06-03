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
	A := nextInt()
	B := nextInt()
	hs := make([]int, N)
	var maxIndex, maxHP int
	for i := 0; i < N; i++ {
		hs[i] = nextInt()
		if hs[i] > maxHP {
			maxIndex = i
			maxHP = hs[i]
		}
	}

	var cnt int
	for {
		cnt++
		isEnd := true
		maxHP = 0
		for i := range hs {
			if i == maxIndex {
				hs[i] -= A
			} else {
				hs[i] -= B
			}
			if hs[i] > maxHP {
				maxIndex = i
				maxHP = hs[i]
			}
			if hs[i] > 0 {
				isEnd = false
			}
		}
		if isEnd {
			break
		}
	}

	fmt.Println(cnt)
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
