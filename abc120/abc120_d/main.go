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
	M := nextInt()
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = nextInt() - 1
		B[i] = nextInt() - 1
	}

	ans := make([]int, M)
	ans[M-1] = N * (N - 1) / 2

	uf := newUnionFind(N)

	for i := M - 1; i >= 1; i-- {
		ans[i-1] = ans[i]
		if uf.root(A[i]) != uf.root(B[i]) {
			ans[i-1] -= uf.size(A[i]) * uf.size(B[i])
			uf.connect(A[i], B[i])
		}
	}

	for i := 0; i < M; i++ {
		fmt.Println(ans[i])
	}
}

type unionFind struct {
	parent []int
}

func newUnionFind(N int) *unionFind {
	parent := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = -1
	}

	return &unionFind{parent: parent}
}

func (uf *unionFind) root(A int) int {
	if uf.parent[A] < 0 {
		return A
	}

	uf.parent[A] = uf.root(uf.parent[A])
	return uf.parent[A]
}

func (uf *unionFind) size(A int) int {
	return uf.parent[uf.root(A)]
}

func (uf *unionFind) connect(A, B int) bool {
	A = uf.root(A)
	B = uf.root(B)
	if A == B {
		return false
	}

	if uf.size(A) < uf.size(B) {
		A, B = B, A
	}

	uf.parent[A] += uf.parent[B]
	uf.parent[B] = A

	return true
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
