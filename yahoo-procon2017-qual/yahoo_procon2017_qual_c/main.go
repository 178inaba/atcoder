package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := nextInt()
	K := nextInt()

	if N == K {
		fmt.Println()
		return
	}

	A := make([]int, 0, K)
	for i := 0; i < K; i++ {
		A = append(A, nextInt())
	}

	S := make([]string, 1, N)
	for i := 0; i < N; i++ {
		S = append(S, nextString())
	}

	targets := make([]string, 0, K)
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for _, a := range A {
		targets = append(targets, S[a])
		S = append(S[:a], S[a+1:]...)
	}

	//sort.Strings(targets)
	base := targets[0]
	for i := 1; i < len(targets); i++ {
		for j := 0; j < min(len(base), len(targets[i])); j++ {
			if base[j] != targets[i][j] {
				base = base[:j]
				break
			}
		}
	}

	for i := 1; i <= len(base); i++ {
		result := true
		for j := 1; j < len(S); j++ {
			if strings.HasPrefix(S[j], base[:i]) {
				result = false
				break
			}
		}

		if result {
			fmt.Println(base[:i])
			return
		}
	}

	fmt.Println(-1)
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

// ---------- Util.
