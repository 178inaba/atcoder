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
	As := make([]int, N)
	for i := 0; i < N; i++ {
		As[i] = nextInt()
	}

	res := make([]int, 10)
	for {
		As = F(As)
		if len(As) == 1 {
			res[As[0]]++
			break
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func F(arr []int) []int {
	x := arr[0]
	y := arr[1]
	arr = append(arr[:0], arr[1:]...)
	arr[0] = (x + y) % 10

	return arr
}

func G(arr []int) []int {
	x := arr[0]
	y := arr[1]
	arr = append(arr[:0], arr[1:]...)
	arr[0] = (x * y) % 10

	return arr
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
