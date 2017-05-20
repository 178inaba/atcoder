package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	H := uint64(nextInt())
	W := uint64(nextInt())
	_, share, mod := areaCalc(H, W, 3)
	if mod == 0 {
		fmt.Println(share)
		return
	}

	var a, b, c, cnt uint64
	for a < share {
		cnt++
		if H < W {
			a += H
		} else {
			a += W
		}
	}
	if H < W {
		W -= cnt
	} else {
		H -= cnt
	}
	area, share, _ := areaCalc(H, W, 2)
	for b < share {
		if H < W {
			b += H
		} else {
			b += W
		}
	}
	c = area - b

	fmt.Println(a - c)
}

func areaCalc(H, W, to uint64) (uint64, uint64, uint64) {
	area := H * W
	share := area / to
	mod := area % to

	return area, share, mod
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
