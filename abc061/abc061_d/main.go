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
	var gs graphs
	for i := 0; i < M; i++ {
		gs = append(gs, graph{a: nextInt(), b: nextInt(), c: nextInt()})
	}

	var max int
	for i, gg := range gs {
		if gg.a == 1 {
			gscp := make(graphs, len(gs))
			copy(gscp, gs)
			gscp[i].isPass = true
			ret := gscp.calc(1, gg.b, gg.c, N)
			if ret > max {
				max = ret
			}
		}
	}

	fmt.Println(max)
}

type graph struct {
	a, b, c int
	isPass  bool
}

type graphs []graph

func (gs *graphs) calc(a, b, c, N int) int {
	var g graph
	var isExist bool
	for i, gg := range *gs {
		if a == gg.a && b == gg.b {
			(*gs)[i].isPass = true
			isExist = true
			g = gg
			break
		}
	}

	if !isExist {
		return 0
	}

	var max int
	for _, gg := range *gs {
		if b == gg.a {
			if gg.isPass {
				fmt.Println("inf")
				os.Exit(0)
			}
			ret := gs.calc(b, gg.b, gg.c, N)
			if ret > max {
				max = ret
			}
		}
	}

	return g.c + max
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
