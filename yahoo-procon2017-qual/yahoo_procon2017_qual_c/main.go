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

	A := make([]int, 0, K)
	for i := 0; i < K; i++ {
		A = append(A, nextInt())
	}

	S := make([]string, 1, N)
	for i := 0; i < N; i++ {
		S = append(S, nextString())
	}

	if N == K {
		fmt.Println()
		return
	}

	targets := make([]string, 0, K)
	for _, a := range A {
		targets = append(targets, S[a])
	}

	sort.Strings(targets)
	base := targets[0]
	charCnt := len(base)
	for i := 0; i < charCnt; i++ {
		result := true
		for j := 1; j < len(targets); j++ {
			if !strings.HasPrefix(targets[j], base) {
				result = false
				break
			}
		}

		if result == true {
			break
		}

		base = base[:len(base)-1]
	}

	if len(base) == 0 {
		fmt.Println(-1)
		return
	}

	for i := 1; i <= N; i++ {
		if !strings.HasPrefix(S[i], base) {
			fmt.Println(base)
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
