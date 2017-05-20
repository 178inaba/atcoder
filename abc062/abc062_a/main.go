package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	x := nextInt()
	y := nextInt()
	groups := [][]int{
		{1, 3, 5, 7, 8, 10, 12},
		{4, 6, 9, 11},
		{2},
	}

	var isSameGroup bool
	for _, group := range groups {
		var isSameGroupX, isSameGroupY bool
		for _, v := range group {
			if x == v {
				isSameGroupX = true
			} else if y == v {
				isSameGroupY = true
			}
		}
		if isSameGroupX && isSameGroupY {
			isSameGroup = true
			break
		}
	}

	if isSameGroup {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.
