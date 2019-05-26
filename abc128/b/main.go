package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type restaurant struct {
	name  string
	point int

	index int
}

type restaurants []restaurant

func (r restaurants) Len() int {
	return len(r)
}

func (r restaurants) Less(i, j int) bool {
	if r[i].name == r[j].name {
		return r[i].point > r[j].point
	}

	return r[i].name < r[j].name
}

func (r restaurants) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	N := nextInt()
	rs := make([]restaurant, N)
	for i := 0; i < N; i++ {
		rs[i] = restaurant{name: nextString(), point: nextInt(), index: i + 1}
	}

	sort.Sort(restaurants(rs))

	for _, r := range rs {
		fmt.Println(r.index)
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
