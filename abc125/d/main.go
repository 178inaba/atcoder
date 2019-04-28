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
	var sum, minus int
	min := int(math.Pow10(9))
	for i := 0; i < N; i++ {
		a := nextInt()
		if a < 0 {
			minus++
		}
		a = abs(a)
		sum += a
		if a < min {
			min = a
		}
	}

	if minus%2 == 0 {
		fmt.Println(sum)
		return
	}

	fmt.Println(sum - min*2)
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

func nextFloat() float64 {
	sc.Scan()
	f64, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}

	return f64
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

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// ---------- Util.
