package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func f(k int) int {
	if k < 3 {
		return math.MaxInt32
	} else if k%3 == 0 {
		return 0
	}
	return 1
}

func main() {
	H := nextInt()
	W := nextInt()
	ans := min(f(H)*W, f(W)*H)
	for i := 1; i < W; i++ {
		a, b, c := H*i, (H/2)*(W-i), (H-H/2)*(W-i)
		ans = min(ans, max(a, max(b, c))-min(a, min(b, c)))
	}
	for i := 1; i < H; i++ {
		a, b, c := W*i, (W/2)*(H-i), (W-W/2)*(H-i)
		ans = min(ans, max(a, max(b, c))-min(a, min(b, c)))
	}
	fmt.Println(ans)
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
