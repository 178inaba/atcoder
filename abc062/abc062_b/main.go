package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	W := nextInt()
	as := make([]string, H)
	for i := 0; i < H; i++ {
		as[i] = nextString()
	}

	s := make([]rune, W+2)
	for i := 0; i < W+2; i++ {
		s[i] = '#'
	}

	fmt.Println(string(s))
	for _, a := range as {
		fmt.Print("#")
		fmt.Print(a)
		fmt.Println("#")
	}
	fmt.Println(string(s))
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
