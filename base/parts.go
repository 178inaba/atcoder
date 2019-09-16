package main

func diff(x, y int) int {
	return abs(-(x) + y)
}

func countRune(s string, r rune) int {
	var cnt int
	for _, sr := range s {
		if sr == r {
			cnt++
		}
	}

	return cnt
}

// Heap. ----------

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

// ---------- Heap.
