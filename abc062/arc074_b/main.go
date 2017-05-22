package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type greaterHeap []int

func (h greaterHeap) Len() int            { return len(h) }
func (h greaterHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h greaterHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *greaterHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *greaterHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type lessHeap greaterHeap

func (h lessHeap) Less(i, j int) bool { return h[i] > h[j] }

func main() {
	N := nextInt()
	as := make([]int, 3*N)
	for i := 0; i < 3*N; i++ {
		as[i] = nextInt()
	}

	h := &greaterHeap{}
	var total int
	for i := 0; i < N; i++ {
		heap.Push(h, as[i])
		total += as[i]
	}
	for i := N; i < 2*N; i++ {
		heap.Push(h, as[i])
		total += as[i]
		total -= heap.Pop(h).(int)
	}
	fmt.Println(h, total)
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
