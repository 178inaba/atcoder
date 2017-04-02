package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type paint struct {
	graph, result map[int]map[int]bool
	colors        []int
}

func main() {
	N := nextInt()
	M := nextInt()

	as, bs := make([]int, M), make([]int, M)
	graph, result := map[int]map[int]bool{}, map[int]map[int]bool{}
	for i := 0; i < M; i++ {
		as[i] = nextInt() - 1
		bs[i] = nextInt() - 1

		if graph[as[i]] == nil {
			graph[as[i]] = map[int]bool{}
		}
		if graph[bs[i]] == nil {
			graph[bs[i]] = map[int]bool{}
		}
		graph[as[i]][bs[i]], graph[bs[i]][as[i]] = true, true
	}
	Q := nextInt()
	vs, ds, cs := make([]int, Q), make([]int, Q), make([]int, Q)
	for i := 0; i < Q; i++ {
		vs[i] = nextInt() - 1
		ds[i] = nextInt()
		cs[i] = nextInt()
	}

	p := paint{graph: graph, result: result, colors: make([]int, N)}

	for i := Q - 1; i >= 0; i-- {
		p.dp(vs[i], ds[i], cs[i])
	}

	for _, c := range p.colors {
		fmt.Println(c)
	}
}

func (p *paint) dp(v, d, c int) {
	if p.result[v] == nil {
		p.result[v] = map[int]bool{}
	}

	if p.result[v][d] {
		return
	}

	p.result[v][d] = true

	if d == 0 {
		p.colors[v] = c
		return
	}

	p.dp(v, d-1, c)

	for u, exist := range p.graph[v] {
		if exist {
			p.dp(u, d-1, c)
		}
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
