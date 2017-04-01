package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type data struct {
	vs, ds, cs []int
	graph      [][]bool
	NColor     []int
	M, Q       int
}

func main() {
	N := nextInt()
	M := nextInt()
	as, bs := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		as[i] = nextInt()
		bs[i] = nextInt()
	}
	Q := nextInt()
	vs, ds, cs := make([]int, Q), make([]int, Q), make([]int, Q)
	for i := 0; i < Q; i++ {
		vs[i] = nextInt()
		ds[i] = nextInt()
		cs[i] = nextInt()
	}

	graph := make([][]bool, M+1)
	for i := 0; i < M+1; i++ {
		graph[i] = make([]bool, M+1)
	}
	for i := 0; i < M; i++ {
		graph[as[i]][bs[i]], graph[bs[i]][as[i]] = true, true
	}

	d := data{vs: vs, ds: ds, cs: cs, graph: graph, NColor: make([]int, N+1), M: M, Q: Q}
	d.recursion(0)

	for _, c := range d.NColor {
		fmt.Println(c)
	}
}

func (d data) recursion(Qi int) {
	if Qi >= d.Q {
		return
	}

	//fmt.Println(Qi, len(d.NColor), len(d.cs))
	d.NColor[d.vs[Qi]] = d.cs[Qi]
	fmt.Println(0)

	for i := 0; i < d.ds[Qi]; i++ {
		for j := 0; j < d.M; j++ {
			if d.graph[d.vs[Qi]+i][i] {
				d.NColor[i] = d.cs[Qi]
			}
		}

		fmt.Println(111)
	}

	d.recursion(Qi + 1)
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
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ---------- Util.
